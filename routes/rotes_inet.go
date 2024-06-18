package routes

import (
	call "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// กำหนด ฺBasicAuth ในการเรียกใช้ API
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		}},
	))

	// URL CALL API
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v3 := api.Group("/v3")
	pungpond := v3.Group("/pungpond")

	v1.Get("/fact/:number", call.Factorial) // 5.1
	pungpond.Get("/", call.ASCII)           // 5.2
	v1.Post("/register", call.ValidTest)    // 6

}
