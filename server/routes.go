package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes() {
	fmt.Printf("Server->registerRoutes I : %+v\n", _cfg.Server.Routes)
	for _, route := range _cfg.Server.Routes {

		if strings.ToUpper(route.Method) == "GET" {
			fmt.Println(route)
			app.Get(route.Path, readRequestQuery, func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"Method": string(c.Request().Header.Method()),
					"Path":   route.Path,
				})
			})
		} else if strings.ToUpper(route.Method) == "POST" {
			fmt.Println(route.Method)
			app.Post(route.Path, func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"Method": string(c.Request().Header.Method()),
					"Path":   route.Path,
				})
			})
		}
	}
}
