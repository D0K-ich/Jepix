package app

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"

	"os"

	//db "Jepix/Server/internal/Database"
	cfg "Jepix/Server/internal/Config"
	"Jepix/Server/internal/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	//"github.com/joho/godotenv"
)

func StartServer() {

	LoadLogger()
	//db.Disconnect()
	//err := (db.СheckHealth())
	//log.Trace(err)

	fmt.Printf(cfg.GetConfDB())

	app := fiber.New(cfg.GetConfigApp())
	log.Trace("Config loaded")

	app.Static("/", "../../../frontend/tmp")
	log.Trace("Static files loaded")

	app.Get("/main", Routes.MainRoute)
	app.Post("/main", Routes.MainRouteP)
	log.Trace("Main page opened")

	app.Get("/authorization", Routes.Authorization)
	app.Post("/authorization", Routes.AuthorizationP)
	app.Post("/reg", Routes.Registration)
	log.Trace("Authorization page opened")

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Listen("127.0.0.1:1111")
}

func LoadLogger() {
	f, err := os.OpenFile("../../logs/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(f)
	log.Trace("Logger loaded")
}
