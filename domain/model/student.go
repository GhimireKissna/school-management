package model

import "time"

type Student struct {
	ID        	uint `json:"id" gorm:"primary_key"`
	CreatedAt 	time.Time
	Class 		string `json:"class"`
	RollNumber 	uint `json:"roll_number"`
	Name    	string `json:"name"`
	Address   	string `json:"address"`
	Email     	string `json:"email"`
	Phone     	int    `json:"phone"`
}