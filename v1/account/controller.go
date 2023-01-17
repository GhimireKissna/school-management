package account

import (
	// "fiber-gorm/database"

	"school_management/config"
	"school_management/domain/model"

	"github.com/gofiber/fiber/v2"
)
func GetGuardians(ctx *fiber.Ctx) error {
	var guardians []model.Guardians
	config.DB.Find(&guardians)
	return ctx.JSON(guardians)
}

func NewGuardian(ctx *fiber.Ctx) error{
	// ctx.Body()
	guardian := new(model.Guardians)
	if err := ctx.BodyParser(&guardian); err != nil{
		return ctx.SendString(err.Error())
	}
	println(guardian.Name,"****************************************")
	config.DB.Create(&guardian)
	return ctx.JSON(guardian)
}
func GetGuardian(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	println(id,"_____________________________________")
	var guardian model.Guardians
	config.DB.Find(&guardian, id)
	return ctx.JSON(guardian)

}
func DeleteGuardian(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var guardian model.Guardians
	config.DB.Delete(&guardian, id)
	return ctx.JSON(guardian)
}
func UpdateGuardian(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	println(id,"***********************************")
	var guardian model.Guardians
	config.DB.Model(&guardian).Where("id = ?", id).Updates(model.Guardians{Name:"sampurna"}).Scan(&guardian)
	// return ctx.JSON(GuardianResponse{
	// 	Name: guardian.Name,
	// 	Address: guardian.Address,
	// })
	return ctx.JSON(guardian)
}

// type GuardianResponse struct {
// 	Name	string `json:"name"`
// 	Address	string `json:"address"`
// }

func ToggleGuardianStatus(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var guardian model.Guardians
	err := config.DB.First(&guardian,id).Error
	if err != nil{
		return ctx.JSON(" guardian not found")
	}
	config.DB.Model(&guardian).Update("is_active", !guardian.IsActive) 
	return ctx.JSON(fiber.Map {
		"message": "User Status Updated sucessfully",
	})
		
	}