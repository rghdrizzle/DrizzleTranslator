package routes

import(
	"github.com/gofiber/fiber/v2"
	"rghdrizzle/drizzletranslator/controller"
)

func UserRoutes(app *fiber.App){
	app.Get("/translate",controller.Translate)
}