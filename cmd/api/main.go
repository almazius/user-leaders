package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	db "github.com/mod"
	"github.com/mod/config"
	"github.com/mod/internal/auth/handlers"
	"log"
)

func main() {
	fmt.Println("hi")

	viperConf, err := config.LoadConfig() // загружаем конфиг для бд из папки config
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}

	db.Connection, err = db.InitPsqlDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	handlers.SetupRoutes(app)
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
