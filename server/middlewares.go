package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func readRequestQuery(c *fiber.Ctx) error {
	//var requestQuery interface{}
	requestQuery := c.Queries()

	log.Printf("Server->MiddleWare->readRequestQuery II : %+v\n", requestQuery)
	return c.Next()
}
