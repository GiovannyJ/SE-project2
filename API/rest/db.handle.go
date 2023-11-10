package rest

import (
	db "API/conn"
	s "API/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type account = s.Account
type updateAcc = s.UpdateAccount
type post = s.Posts
type updatePost = s.UpdatePosts
type images = s.Images
type updateImages = s.UpdateImages
type loginData = s.LogIn

/*
*=================GET METHOD HANDLERS==================
*/

/*
*TESTED WORKING
GETS all accounts
can query by id, fname, lname, fullname, email, username, and order
*/
func GetAccounts(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})

	queryParams := map[string]string{
		"id":       c.Query("id"),
		"fname":    c.Query("fname"),
		"lname":    c.Query("lname"),
		"fullname": c.Query("fullname"),
		"email":    c.Query("email"),
		"username": c.Query("username"),
		"order":	c.Query("order"),
	}
	
	queryKey := ""
	for key, value := range queryParams {
		if len(value) > 0 {
			queryKey = key
			break
		}
	}
	
	if queryKey != "" {
		query[queryKey] = queryParams[queryKey]
	}

	c.IndentedJSON(http.StatusOK, db.Accounts_GET(query))
}

/*
*TESTED WORKING
GETS all posts:
can query by id, title, genre, authorID, numUp, numDown, postedDate, order
*/
func GetPosts(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})

	queryParams := map[string]string{
		"id":     		c.Query("id"),
		"title":    	c.Query("title"),
		"genre": 		c.Query("genre"),
		"authorId":		c.Query("authorID"),
		"numUp":   		c.Query("numUp"),
		"numDown":  	c.Query("numDown"),
		"postedDate": 	c.Query("date"),
		"order":		c.Query("order"),
	}
	
	queryKey := ""
	for key, value := range queryParams {
		if len(value) > 0 {
			queryKey = key
			break
		}
	}
	
	if queryKey != "" {
		query[queryKey] = queryParams[queryKey]
	}

	c.IndentedJSON(http.StatusOK, db.Posts_GET(query))
}

/*
*TESTED WORKING
GETS all posts with author info as well as image info
can query by, postID, authorID, fullname, username, title, numUp,
numDown, genre, date, order
*/
func GetPostsFullContext(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})

	queryParams := map[string]string{
		"p.id":     	c.Query("postID"),
		"a.id":			c.Query("authorID"),
		"fullname": 	c.Query("fullname"),
		"username": 	c.Query("username"),
		"title":    	c.Query("title"),
		"numupq":   	c.Query("numUp"),
		"numdown":  	c.Query("numDown"),
		"genre": 		c.Query("genre"),
		"postedDate": 	c.Query("date"),
		"order":		c.Query("order"),
	}
	
	queryKey := ""
	for key, value := range queryParams {
		if len(value) > 0 {
			queryKey = key
			break
		}
	}
	
	if queryKey != "" {
		query[queryKey] = queryParams[queryKey]
	}

	results, err := db.PostsFullContext_GET(query)
	if err != nil{
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

/*
*=================POST METHOD HANDLERS==================
*/

/*
*TESTED WORKING
Creates new account
Request body shaped like account struct without id and access level
*/
func NewAccount(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	
	var newAcc account

	if err := c.BindJSON(&newAcc); err != nil{
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := db.CreateNewAccount(newAcc)
	if err != nil{
		if _, ok := err.(*s.AccountExistsError); ok{
			c.IndentedJSON(http.StatusBadRequest, "Account Already Exits")
		}else{
			c.IndentedJSON(http.StatusInternalServerError, nil)
		}
		return
	}
	q := map[string]interface{}{
		"username": newAcc.Username,
	}

	c.IndentedJSON(http.StatusCreated, db.Accounts_GET(q))
}

/*
*TESTED WORKING
CREATES new post
Request body shaped like posts struct
*/
func NewPost(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	
	var newPost post

	if err := c.BindJSON(&newPost); err != nil{
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := db.CreateNewPost(newPost)
	if err != nil{
		if _, ok := err.(*s.AccountExistsError); ok{
			c.IndentedJSON(http.StatusBadRequest, "Account Already Exits")
		}else{
			c.IndentedJSON(http.StatusInternalServerError, nil)
		}
		return
	}
	
	c.IndentedJSON(http.StatusCreated, "post created")
}

/*
*TESTED WORKING
Allows user to login using email or username
Request body shaped like login struct: must have empty string for username or email if not used
returns limited version of account struct
Note: if username and password as passed as guest then its guest login
*/
func Login(c *gin.Context) {
    var login loginData

    if err := c.BindJSON(&login); err != nil {
        c.IndentedJSON(http.StatusBadRequest, "cannot bind")
        return
    }

    logicCheck := func(param *string) bool {
        return param != nil && *param == ""
    }

    if (logicCheck(login.Username) || logicCheck(login.Email)) && logicCheck(login.Password) {
        c.IndentedJSON(http.StatusBadRequest, "fail logic check")
        return
    }

    if login.Username != nil && *login.Username == "guest" && login.Password != nil && *login.Password == "guest"{
        result := db.GuestLogin()
        c.IndentedJSON(http.StatusOK, result)
        return
    }

    data, err := db.GetLoginInfo(login.Username, login.Password, login.Email)

    if err != nil {
        if err.Error() == "Invalid password" {
            c.IndentedJSON(http.StatusUnauthorized, "invalid password")
        } else {
            c.IndentedJSON(http.StatusNotFound, "invalid password")
        }
        return
    }

    if data != nil {
        userQ := map[string]interface{}{
            "id": data.ID,
        }
        userDetails := db.Accounts_GET(userQ)
        c.IndentedJSON(http.StatusOK, userDetails)
    } else {
        c.IndentedJSON(http.StatusNotFound, nil)
    }
}



/*
*=================PATCH METHOD HANDLERS==================
*/
// func UpdateAcc(c *gin.Context) {
//     c.Header("Access-Control-Allow-Origin", "*")
    
//     var upACC updateAcc

//     if err := c.BindJSON(&upACC); err != nil {
//         c.IndentedJSON(http.StatusBadRequest, nil)
//         return 
//     }

// 	err := db.UpdateData(upACC.Old, upACC.New)

//     if err != nil {
// 		if _, ok := err.(*s.UpdateNotCompleteError); ok{
// 			c.IndentedJSON(http.StatusFailedDependency, nil)
// 		}else{
// 			c.IndentedJSON(http.StatusInternalServerError, nil)
// 		}
//         return 
// 	}
//     c.IndentedJSON(http.StatusOK, upACC.New)
// }
