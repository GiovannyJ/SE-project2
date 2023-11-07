package models

import "fmt"

type Account struct {
	ID       int    `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Pwd 	 string `json:"pwd"`
	Pnum     int    `json:"pnum"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}

type UpdateAccount struct {
	Old Account `json:"old"`
	New Account `json:"new"`
}

type AccountExistsError struct {
	Username string
}
func (e *AccountExistsError) Error() string {
	return fmt.Sprintf("Account with username %s already exists", e.Username)
}
type AccountNotExistsError struct {
	Msg string
}
func (e *AccountNotExistsError) Error() string {
	return fmt.Sprintf(e.Msg)
}

type UpdateNotCompleteError struct{
	Msg string
}
func (e *UpdateNotCompleteError) Error() string{
	return fmt.Sprintf(e.Msg)
}