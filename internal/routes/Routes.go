package Routes

//todo Сделать отдельную авторизацию для админа. Реализовать регистрацию.
//todo Сделать передачу cookie-флагов при авторизации.
//todo Сделать генерацию секретного слова для криптографии
import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func MainRoute(c *fiber.Ctx) error {
	return c.Render("main", fiber.Map{})
}

func MainRouteP(c *fiber.Ctx) error {
	return c.JSON("321232132")
}

func Authorization(c *fiber.Ctx) error {
	return c.Render("index3", fiber.Map{})
}

func AuthorizationP(c *fiber.Ctx) error {

	user := c.FormValue("login")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user == "Dok" && pass == "Dok" {
		return c.SendStatus(fiber.StatusUnauthorized) // возвращать страничку админа
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}

func Registration(c *fiber.Ctx) error {

	//todo Прикрутить модельку юзера; сделать заполнение бд.
	// todo Проанализировать, как будет лучше, криптовать данные логина или нет

	user := c.FormValue("login")
	pass := c.FormValue("pass")
	email := c.FormValue("email")
	phone := c.FormValue("phone")

	claims := jwt.MapClaims{
		"name":     user,
		"email":    email,
		"phone":    phone,
		"password": pass,
		"admin":    false,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON("Not available for registration")
	}

	//
	return c.JSON("Registered!%s", t)
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
