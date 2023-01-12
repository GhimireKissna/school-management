package model
import "time"

type Guardians struct{
	ID        		uint `json:"id" gorm:"primary_key"`
	CreatedAt 		time.Time
	Name    		string `json:"name"`
	Address   		string `json:"address"`
	Email     		string `json:"email"`
	Phone     		int    `json:"phone"`
	StudentRefer 	int `json:"student_id"`
	Student 		Student `gorm:"foreignkey:StudentRefer"` 
}