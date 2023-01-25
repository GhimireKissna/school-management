package account

import (
	"fmt"
	"school_management/config"
	"school_management/domain/model"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	)
func CheckIn(ctx *fiber.Ctx) error{
	var student model.Student
	var attendance model.Attendance
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	err := config.DB.First(&student,idInt).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
		"message":"Student Was Not Found",
		})
	}

	y,m,d:= time.Now().Date()
	str1 := strconv.Itoa(y)
	str2 := strconv.Itoa(int(m))
	str3 := strconv.Itoa(d)
  	year := str1
    month:= str2
	day  := str3
	err = config.DB.Debug().Where("student=? AND year =? AND month =? AND day =?", id, year,month,day).First(&attendance).Error
		if err == nil{
			return ctx.JSON(fiber.Map{
				"error":"User already CheckIn",
		})
	}
	var checkIn string 
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

    str4 := strconv.Itoa(hour)
    str5 := strconv.Itoa(minute)
    str6 := strconv.Itoa(second)
    checkIn = str4 + ":" + str5 + ":" + str6 
	
	query := fmt.Sprintf("INSERT INTO attendances(year,month,day,check_in,student) VALUES ('%s','%s','%s', '%v', %d)", year,month,day, checkIn, idInt)
	config.DB.Exec(query)
	
	return ctx.JSON(fiber.Map{
		"Message":"Student Check In Sucessfully",
	})
}

func CheckOut(ctx *fiber.Ctx) error{
	var student model.Student
	var attendance model.Attendance
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	err := config.DB.First(&student,idInt).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
		"message":"Student Was Not Found",
		})
	}
	var checkOut string
	y,m,d := time.Now().Date()
	str4 := strconv.Itoa(y)
	str5 := strconv.Itoa(int(m))
	str6 := strconv.Itoa(d)
	year := str4
	month := str5
	day := str6
	err = config.DB.Debug().Where("student =? AND year =? AND month =? AND day =?",id,year,month,day).First(&attendance).Error
		if err != nil {
		return ctx.JSON(fiber.Map{
			"error":"Student Was Not Check In",
		})
	 }
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	str1:= strconv.Itoa(hour)
	str2 := strconv.Itoa(minute)
	str3:= strconv.Itoa(second)
	checkOut = str1 + ":" + str2 + ":" +str3
	query := fmt.Sprintf("UPDATE attendances SET check_out = '%s' WHERE student = %d AND year= '%s' AND month = '%s' AND day ='%s'",checkOut,idInt,year,month,day)
	config.DB.Exec(query)
	return ctx.JSON(fiber.Map{
		"message":"Student Check Out Sucessfully",
	})
}
