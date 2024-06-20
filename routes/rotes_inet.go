package routes

import (
	call "go-fiber-test/controllers" //5.3

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// กำหนด ฺBasicAuth ในการเรียกใช้ API
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566", // 5.0
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

	// * Dogs
	dog := v1.Group("/dog")
	dog.Get("", call.GetDogs)
	dog.Get("/filter", call.GetDog)
	dog.Get("/json", call.GetDogsJson) // 7.2
	dog.Post("/", call.AddDog)
	dog.Put("/:id", call.UpdateDog)
	dog.Delete("/:id", call.RemoveDog)
	dog.Get("/del", call.GetDogsDel) // 7.0.2
	dog.Get("/scope", call.Getscope) // 7.1

	// * company 7.0.2
	company := v1.Group("/company")
	company.Get("", call.Getcompanys)
	company.Post("/", call.Addcompany)
	company.Put("/:id", call.Updatecompany)
	company.Delete("/:id", call.Removecompany)

}
