package account
import "github.com/gofiber/fiber/v2"

func  AccountRoutes(app *fiber.App){
	app.Get("/api/v1/guardians/list/", GetGuardians)
	app.Post("/api/v1/guardians/create/", NewGuardian)
	app.Get("/api/v1/guardians/:id/get/", GetGuardian)
	app.Delete("/api/v1/guardians/:id/delete/", DeleteGuardian)
	app.Patch("/api/v1/guardians/:id/update/", UpdateGuardian)
	app.Patch("/api/v1/guardians/:id/toggle-status/", ToggleGuardianStatus)
	app.Patch("/api/v1/guardians/:id/change/", ChangePassword)

	app.Get("/api/v1/principle/list/", GetPrinciple)
	app.Post("/api/v1/principle/create/", NewPrinciple)
	app.Get("/api/v1/principle/:id/get/", GetPrincipleById)
	app.Delete("/api/v1/principle/:id/delete/", DeletePrinciple)
	app.Patch("/api/v1/principle/:id/update/", UpdatePrinciple)
	app.Patch("/api/v1/principle/:id/toggle-status/", TogglePrincipleStatus)

	app.Get("/api/v1/students/list/", GetStudents)
	app.Post("/api/v1/student/create/", NewStudents)
	app.Get("/api/v1/student/:id/get/", GetStudent)
	app.Delete("/api/v1/student/:id/delete/", DeleteStudent)
	app.Patch("/api/v1/student/:id/update/", UpdateStudent)
	app.Patch("/api/v1/student/:id/toggle-status/", ToggleStudentStatus)


	app.Post("api/v1/student/:id/checkin/", CheckIn)
	app.Patch("api/v1/student/:id/checkout/", CheckOut)
}