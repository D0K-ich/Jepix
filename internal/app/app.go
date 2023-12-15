package app

import (
	jwtware "github.com/gofiber/contrib/jwt"

	"os"

	cfg "Jepix/internal/Config"
	db "Jepix/internal/Database"
	"Jepix/internal/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	//"github.com/joho/godotenv"
)

func StartServer() {

	//LOGGER
	LoadLogger()

	//DB
	DbStatus := db.CheackHealth()
	log.Trace(DbStatus)

	//Init app
	app := fiber.New(cfg.GetConfigApp())
	log.Trace("App with config loaded")

	//Load TMP files
	app.Static("/", "../../frontend/tmp")
	log.Trace("Static files loaded")

	//Middlewares for main page
	app.Get("/main", Routes.MainRoute)

	//Middlewares for authorization page
	app.Get("/authorization", Routes.Authorization)
	app.Post("/authorization/data", Routes.AuthorizationData)

	//Middlewares for authorization page
	app.Get("/registration", Routes.Authorization)
	app.Post("/registration/data", Routes.Registration)

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
