package main

import (
	"github.com/AnggaDanarP/Online-Learning-Plathform/database"
	"github.com/AnggaDanarP/Online-Learning-Plathform/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // to get the cookie that we send and send it back
	})) // Enable CORS to frondend

	routes.Setup(app)

	app.Listen(":8000")
}
