package models

type Account struct {
	ID       int    `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
	Pnum     int    `json:"pnum"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}