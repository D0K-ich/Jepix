package Config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

import (
	models "Jepix/Server/internal/Models"
	"fmt"
	"os"
)

var con = godotenv.Load("../../.env")

func GetConfDB() string {
	Config := models.ConfigDB{
		User:    os.Getenv("User"),
		Pass:    os.Getenv("Pass"),
		Dbname:  os.Getenv("Dbname"),
		HostDB:  os.Getenv("HostDB"),
		Port:    os.Getenv("Port"),
		Sslmode: os.Getenv("Sslmode"),
	}

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", Config.User, Config.Pass, Config.Dbname, Config.HostDB, Config.Port, Config.Sslmode)
}

func GetConfigApp() fiber.Config {
	engine := html.New("../../../frontend", ".html")

	return fiber.Config{
		AppName:       "Jepix",
		CaseSensitive: true, // разница в carmelCase в запросах
		//BodyLimit:          5,    // лимит запросов
		//Concurrency:        2,    // параллельные подключение
		EnableIPValidation: true, // проверка ip, первый рабочий будет возвращаться
		Views:              engine,
	}
}
