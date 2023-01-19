package course
import (
	"school_management/config"
	"school_management/domain/model"
	"github.com/gofiber/fiber/v2"
)
func GetCourse(ctx *fiber.Ctx) error{
	var Courses []model.Courses
	config.DB.Find(&Courses)
	return ctx.JSON(Courses)
}
func NewCourse(ctx *fiber.Ctx) error{
	course := new(model.Courses)
	if err := ctx.BodyParser(&course); err != nil{
		return ctx.SendString(err.Error())
	}
	config.DB.Create(&course)
	return ctx.JSON(course)
}
func GetCourseById(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var course model.Courses
	config.DB.Find(&course,id)
	return ctx.JSON(course)
}
func DeleteCourse(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var course model.Courses
	config.DB.Delete(&course, id)
	return ctx.JSON(course)
}
func UpdateCourse(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var course model.Courses
	config.DB.Model(&course).Where("id =?",id).Updates(model.Courses{Name: "php"}).Scan(&course)
	return ctx.JSON(course)
}
func ToggleCourseStatus(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var course model.Courses
	err := config.DB.First(&course, id).Error
	if err != nil {
		return ctx.JSON("Course Not Found")
	}
	config.DB.Model(&course).Update("is_active", !course.IsActive)
	return ctx.JSON(fiber.Map{
		"message":"Courses Updated Sucessfully",
	})
}