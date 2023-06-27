package server

import (
	"fmt"
	"log"

	db "github.com/degreane/ezekod.com/database"
	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

func Start() {
	app = fiber.New(readConfig())
	db.Initialize(readDBConfig())
	registerRoutes()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", _cfg.Server.Config.Host, _cfg.Server.Config.Port)))

}
