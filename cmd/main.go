package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"main.go/pkg/repository"
	"main.go/pkg/service"
	"main.go/pkg/telegram"
	"os"
)

var bot *tgbotapi.BotAPI

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("init config err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("init db err: %s", err.Error())
	}

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_API"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Debug = true

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	tgBot := telegram.NewBot(bot, services)
	err = tgBot.Start()
	if err != nil {
		logrus.Fatalf("bot.start failed: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
