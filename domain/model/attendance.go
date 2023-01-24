package model
type Attendance struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Student   int    `json:"student"`
	CheckIn   string `json:"checkin"`
	CheckOut  string `json:"checkout"`
	Day       string `json:"day"`
	Month   string `json:"month"`
	Year      string `json:"year"`
}