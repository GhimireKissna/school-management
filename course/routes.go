package course
import (
	"github.com/gofiber/fiber/v2"
)
func CourseRoutes(app *fiber.App){
	app.Get("/api/course/list/", GetCourse)
	app.Post("/api/course/create/", NewCourse)
	app.Get("/api/course/:id/get/", GetCourseById)
	app.Delete("/api/course/:id/delete/", DeleteCourse)
	app.Patch("/api/course/:id/update/", UpdateCourse)
	app.Patch("/api/course/:id/toggle-status/", ToggleCourseStatus)
}