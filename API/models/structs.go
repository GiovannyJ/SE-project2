package models

import "fmt"

type Account struct {
	ID       	int    	`json:"id"`
	Fname    	string 	`json:"fname,omitempty"`
	Lname    	string 	`json:"lname,omitempty"`
	Fullname 	string 	`json:"fullname,omitempty"`
	Email    	string 	`json:"email,omitempty"`
	Pwd 	 	string 	`json:"pwd,omitempty"`
	Pnum     	int    	`json:"pnum,omitempty"`
	Username 	string 	`json:"username"`
	Accesslvl 	string	`json:"acesslevel"`
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


type Images struct{
	ID 		int `json:"id"`
	ImgName string `json:"imgname"`
}

type UpdateImages struct{
	Old Images `json:"old"`
	New Images `json:"new"`
}

type Posts struct{
	ID 			int 	`json:"id"`
	Title 		string  `json:"title"`
	Descr 		string  `json:"descr"`
	Genre 		string  `json:"genre"`
	AuthorID 	int 	`json:"authorId"`
	NumUp 		int 	`json:"numUp"`
	NumDown 	int 	`json:"numDown"`
	PicID 		int 	`json:"picId,omitempty"`
	PostedDate 	string		`json:"postedDate"`
}

type UpdatePosts struct{
	Old Posts `json:"old"`
	New Posts `json:"new"`
}

type FullContextPost struct{
	ID 			int 	`json:"id"`
	Title 		string 	`json:"title"`
	Descr 		string 	`json:"descr"`
	Genre 		string 	`json:"genre"`
	AuthorInfo 	Account `json:"authorInfo"`
	NumUp 		int 	`json:"numUp"`
	NumDown 	int 	`json:"numDown"`
	ImageInfo 	Images 	`json:"picInfo"`
	PostedDate 	string		`json:"postedDate"`
}

type LogIn struct {
    Username *string `json:"username"`
    Email 	 *string `json:"email"`
	Password *string `json:"password"`
}