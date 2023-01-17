package account
import "github.com/gofiber/fiber/v2"

func AccountRoutes(app *fiber.App){
	app.Get("/api/v1/guardians/list/", GetGuardians)
	app.Post("/api/v1/guardians/create/", NewGuardian)
	app.Get("/api/v1/guardians/:id/get/", GetGuardian)
	app.Delete("/api/v1/guardians/:id/delete/", DeleteGuardian)
	app.Patch("/api/v1/guardians/:id/update/", UpdateGuardian)
	app.Patch("/api/v1/guardians/:id/toggle-status/", ToggleGuardianStatus)
}