package model

import (
	"time"
)

type Principle struct{
	ID        	uint `json:"id" gorm:"primary_key"`
	CreatedAt 	time.Time
	Name    	string `json:"name"`
	Address   	string `json:"address"`
	Email     	string `json:"email"`
	Phone     	int    `json:"phone"`
	IsActive	bool   `json:"isactive"`

}
