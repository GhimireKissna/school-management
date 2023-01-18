package account
import (
	"school_management/config"
	"school_management/domain/model"
	"github.com/gofiber/fiber/v2"
)
func GetStudents(ctx *fiber.Ctx) error{
	var student []model.Student
	config.DB.Find(&student)
	return ctx.JSON(student)
}

func NewStudents(ctx *fiber.Ctx) error{
	students := new(model.Student)
	if err := ctx.BodyParser(&students); err != nil{
		return ctx.SendString(err.Error())
	}
		config.DB.Create(&students)
		return ctx.JSON(students)
}
func GetStudent(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var student model.Student
	config.DB.Find(&student,id)
	return ctx.JSON(student)
}
func DeleteStudent(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var student model.Student
	config.DB.Delete(&student, id)
	return ctx.JSON(student)
}
func UpdateStudent(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var student model.Student
	config.DB.Model(student).Where("id = ?", id).Updates(model.Student{Name:"sampurna"}).Scan(&student)
	return ctx.JSON(student)
}
func ToggleStudentStatus(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var student model.Student
	err := config.DB.First(&student,id).Error
	if err != nil{
		return ctx.JSON("Student Not Found")
	}
	config.DB.Model(&student).Update("is_active",!student.IsActive)
	return ctx.JSON(fiber.Map{
		"message":"Student Updated Sucessfully",
	})
		
}
