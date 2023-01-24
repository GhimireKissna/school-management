package account

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"school_management/config"
	"school_management/domain/model"
)
func GetPrinciple(ctx *fiber.Ctx) error{
	var principle []model.Principle
	config.DB.Find(&principle)
	return ctx.JSON(principle)
}
func NewPrinciple(ctx *fiber.Ctx) error{
	principle := new(model.Principle)
	if err := ctx.BodyParser(&principle); err != nil{
		return ctx.SendString(err.Error())
	}
	config.DB.Create(&principle)
	return ctx.JSON(principle)

}
func GetPrincipleById(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var principle model.Principle
	config.DB.Find(&principle,id)
	return ctx.JSON(principle)
}

func DeletePrinciple(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var principle model.Principle
	config.DB.Delete(&principle,id)
	return ctx.JSON(principle)
}
func UpdatePrinciple(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var principle model.Principle
	config.DB.Model(principle).Where("id = ?", id).Updates(model.Principle{Name:"sampurna"}).Scan(&principle)
	return ctx.JSON(principle)
}
func TogglePrincipleStatus(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var principle model.Principle
	err := config.DB.First(&principle,id).Error
	fmt.Println(principle,"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	if err != nil {
		return ctx.JSON("Principle Not Found")
	}
	config.DB.Model(&principle).Update("is_active", !principle.IsActive)
	return ctx.JSON(fiber.Map{
		"message":"Principle Status Updated Sucessfully",
	})
}