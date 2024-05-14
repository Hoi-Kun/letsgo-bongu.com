package middleware

import (
	"net/http"

	"github.com/Hoi-Kun/letsgo-bongu.com/menu"
	"github.com/gofiber/fiber/v2"
)

func BindMenu(c *fiber.Ctx) error {
	m := fiber.Map{}
	m["Whiskies"] = menu.Whiskies
	m["TCs"] = menu.TCs
	m["Foods"] = menu.Foods
	m["Menu"] = menu.Menus
	if err := c.Bind(fiber.Map{"Menu": m}); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Next()
}
