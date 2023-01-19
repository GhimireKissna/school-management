package main
import (
	"school_management/domain/model"
	"school_management/config"
)

func main(){
	db, err := config.Connect()
	if err != nil {
		println("Error Connecting to database")
	}
	db.AutoMigrate(&model.Guardians{},&model.Principle{}, &model.Student{})
	db.AutoMigrate(&model.Courses{})
	println("Migartion successful.")
}