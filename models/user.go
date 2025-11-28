package models

type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm-password"`
}
