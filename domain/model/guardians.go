package model
import "time"

type Guardians struct{
	ID        		uint `json:"id" gorm:"primary_key"`
	CreatedAt 		time.Time
	Name    		string `json:"name"`
	Address   		string `json:"address"`
	Phone			int `json:"phone"`
	Email     		string `json:"email"`
	Password		string	`json:"password"`
	IsActive		bool `json:"isactive"`
	// StudentRefer 	int `json:"-"`
	// Student 		Student `gorm:"foreignkey:StudentRefer"` 
}