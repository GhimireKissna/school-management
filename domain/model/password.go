package model

type ChangePassword struct{
	OldPassword		string `json:"old_password"`
	NewPassword		string	`json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`

}