package models

import (
	"fmt"
	"time"
)

type Account struct {
	ID       	int    	`json:"id"`
	Fname    	string 	`json:"fname,omitempty"`
	Lname    	string 	`json:"lname,omitempty"`
	Fullname 	string 	`json:"fullname,omitempty"`
	Email    	string 	`json:"email,omitempty"`
	Pwd 	 	string 	`json:"pwd,omitempty"`
	Pnum     	int    	`json:"pnum,omitempty"`
	Username 	string 	`json:"username"`
	Accesslvl 	string	`json:"accesslevel"`
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
	Size 	string `json:"size"`
	Date 	string `json:"date"`
}

type FileUpload struct {
	Name *string    `json:"name"`
	Size *string     `json:"size"`
	Date *time.Time `json:"date"`
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
	PostedDate 	string	`json:"postedDate"`
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
	PostedDate 	string	`json:"postedDate"`
}

type LogIn struct {
    Username *string `json:"username"`
    Email 	 *string `json:"email"`
	Password *string `json:"password"`
}

type Comment struct{
	ID 			int 	`json:"id"`
	PostID		int		`json:"postID"`
	AuthorID 	int 	`json:"authorID"`
	Content		string 	`json:"content"`
	NumUp 		int 	`json:"numUp"`
	NumDown 	int 	`json:"numDown"`
	PostedDate 	string	`json:"postedDate"`
}

type CommentFullContext struct{
	CommentInfo Comment `json:"commentInfo"`
	CommenterInfo Account `json:"commenterInfo"`
	PostInfo Posts `json:"postInfo"`
	PostAuthorInfo Account `json:"postAuthorInfo"`
	ImageInfo Images `json:"imageInfo"`
}

type UpdateComment struct{
	Old Comment `json:"old"`
	New Comment `json:"new"`
}

type JSONData struct{
	Data interface{} `json:"data"`
}

