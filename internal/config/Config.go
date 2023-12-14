package Config

import (
	models "Jepix/internal/Models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load("../../internal/config/.env")

func GetConfDB() string {
	Config := models.ConfigDB{
		User:       os.Getenv("LOGINDB"),
		Pass:       os.Getenv("PASSDB"),
		Collection: os.Getenv("COLLECTIONDB"),
	}
	return fmt.Sprintf("mongodb+srv://%s:%s@%s.yshipgd.mongodb.net/", Config.User, Config.Pass, Config.Collection)
}

func GetConfigApp() fiber.Config {
	engine := html.New("../../frontend", ".html")

	return fiber.Config{
		AppName:       "Jepix",
		CaseSensitive: true, // разница в carmelCase в запросах
		//BodyLimit:          5,    // лимит запросов
		//Concurrency:        2,    // параллельные подключение
		EnableIPValidation: true, // проверка ip, первый рабочий будет возвращаться
		Views:              engine,
	}
}
