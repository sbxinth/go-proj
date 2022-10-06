package routes

import (
	ctrl "go-proj/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Router(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v1.Use(basicauth.New(basicauth.Config{Users: map[string]string{
		"testgo": "6102565",
	}}))
	v1.Get("/test", ctrl.SendHi)
	v1.Post("/User", ctrl.UserADD)
	v2.Get("/Data", ctrl.GetGEN)
	v2.Get("/DataP", ctrl.GetParm)
}
