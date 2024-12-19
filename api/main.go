// declare that files belongs to main package, tells go its executable file and not a library, so it should have a main() function, entry point of the program
package main

import (
	"go-api/internal/web3"
)

func main() {

	web3.Connect()
	// if err := godotenv.Load(); err != nil {
    //     log.Fatal("Error loading .env file")
    // }

    // // Connect to database
    // database.ConnectDB()

	
	// app := fiber.New()
	// handlers.SetupPostRoutes(app)
	// // app.Use(middleware.lo)
	// log.Fatal(app.Listen(":3000")) //App listen starts the server, log.Fatal logs the error if any
}