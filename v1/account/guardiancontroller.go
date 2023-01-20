package account

import (
	"school_management/config"
	"school_management/domain/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)
func hashPassword(password string)(string, error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}
func checkPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}

func GetGuardians(ctx *fiber.Ctx) error {
	var guardians []model.Guardians
	config.DB.Find(&guardians)
	return ctx.JSON(guardians)
}

func NewGuardian(ctx *fiber.Ctx) error{
	guardian := new(model.Guardians)
	if err := ctx.BodyParser(&guardian); err != nil{
		return ctx.SendString(err.Error())
	}
	hash, err := hashPassword(guardian.Password)
	if err != nil{
		return ctx.JSON(hash)
	}
	guardian.Password = hash
	if err := config.DB.Create(&guardian).Error; err != nil{
		return ctx.JSON(hash)
	}
	// config.DB.Create(&guardian)
	return ctx.JSON(guardian)
}
func GetGuardian(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
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
		"message": "Guardians Status Updated sucessfully",
	})
		
	}
func ChangePassword(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var guardian model.Guardians
	var changePassword model.ChangePassword
		err := config.DB.First(&guardian,id).Error
	if err != nil{
		return ctx.JSON(fiber.Map{
			"message":"Guardian doesnot exist",
		})
	}
	if err := ctx.BodyParser(&changePassword); err != nil{
		return ctx.JSON(fiber.Map{
			"message":err,
		})
	}
	if !checkPasswordHash(changePassword.OldPassword, guardian.Password){
		return ctx.JSON(fiber.Map{
			"message":"invalid password",
		})
	}
		
		newHash, err := hashPassword(changePassword.NewPassword)
		if err != nil{
			return ctx.JSON(fiber.Map{
				"message": err,
		})
	}
	if err := config.DB.Model(&guardian).Update("password",newHash).Error; err != nil{
		return ctx.JSON(fiber.Map{
			"message": err,
		})
	}
		return ctx.JSON(fiber.Map{
			"message":"Password Change Sucessfully",
		})
}