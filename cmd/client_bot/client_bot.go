package main

import (
	"log"
	"os"

	scheduler "github.com/Ksenofontovas/solution_cup/domain"
	"github.com/Ksenofontovas/solution_cup/internal/UI/tgbot"
	"github.com/Ksenofontovas/solution_cup/internal/repository"
	"github.com/Ksenofontovas/solution_cup/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	db.AutoMigrate(&scheduler.Task{})
	if err != nil {
		log.Fatalf("failed to migrate: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)

	botDebug := false
	updateTimeout := 60

	bot, err := tgbot.NewTgBot(os.Getenv("BOT_API_KEY"), botDebug, updateTimeout, service)
	if err != nil {
		log.Panic(err)
	}

	bot.GetUpdates()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
