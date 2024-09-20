package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"homework_bot/internal/application/services"
	"homework_bot/internal/bot/telegram"
	"homework_bot/internal/infrastructure/configs"
	repository "homework_bot/internal/infrastructure/repositories"
	"os"
)

var bot *tgbotapi.BotAPI

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("init configs err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %s", err.Error())
	}

	db, err := configs.NewPostgresDB(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
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
	service := services.NewService(repos)

	tgBot := telegram.NewBot(bot, service)
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
