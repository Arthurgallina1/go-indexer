// declare that files belongs to main package, tells go its executable file and not a library, so it should have a main() function, entry point of the program
package main

import (
	"go-api/internal/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	handlers.SetupPostRoutes(app)
	// app.Use(middleware.lo)
	log.Fatal(app.Listen(":3000")) //App listen starts the server, log.Fatal logs the error if any
}