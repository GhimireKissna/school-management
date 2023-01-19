package model
import "time"
type Courses struct{
	ID        	uint `json:"id" gorm:"primary_key"`
	CreatedAt 	time.Time
	Name		string `json:"name"`
	Discription	string `json:"discription"`
	IsActive bool		`json:"isactive"`
}