package main

import (
	"context"
	"github.com/Craxe99/chat"
	"github.com/Craxe99/chat/pkg/handler"
	"github.com/Craxe99/chat/pkg/repository"
	"github.com/Craxe99/chat/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Установка формата лога
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация viper для чтения файла config.yml
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occurred while reading config: %s", err.Error())
	}

	// Подключение к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error occurred while connecting to DB: %s", err.Error())
	}

	// Создание экземпляра репозитория
	repos := repository.NewRepository(db)
	// Создание экземпляра сервисов
	services := service.NewService(repos)
	// Создание экземпляра обработчиков
	handlers := handler.NewHandler(services)

	// Создание экземпляра сервера
	srv := new(chat.Server)
	// Запуск сервера в горутине
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error occurred while running httpServer: %s", err.Error())
		}
	}()

	logrus.Print("Chat Started")

	// Создание канала, отслеживающего завершение сервиса
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Chat Shutting Down")

	// Отключение сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured while server shutting down: %s", err.Error())
	}
	// Отключение от базы данных
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured while server closing db: %s", err.Error())
	}
}

// Инициализация viper
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
