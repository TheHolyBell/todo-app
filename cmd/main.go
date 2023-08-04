package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"rest_api"
	"rest_api/pkg/handler"
	"rest_api/pkg/repository"
	"rest_api/pkg/service"
	"syscall"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing configs: %v\n", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading environment variables: %v\n", err)
	}

	cfgMap := viper.GetStringMapString("db")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfgMap["host"],
		Port:     cfgMap["port"],
		Username: cfgMap["username"],
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfgMap["name"],
		SSLMode:  cfgMap["sslmode"],
	})

	if err != nil {
		logrus.Fatalf("Could not connect to db: %v\n", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	srv := new(rest_api.Server)
	go func() {
		if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running server: %v\n", err)
		}
	}()

	logrus.Printf("TodoApp has started on port %s...", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp is shutting down...")
	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutdown: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occurred while closing database connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
