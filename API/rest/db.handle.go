package rest

import (
	db "API/conn"
	s "API/models"
	"fmt"

	// "strings"

	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)


type account = s.Account
type updateAcc = s.UpdateAccount
type post = s.Posts
type updatePost = s.UpdatePosts
type images = s.Images
type updateImages = s.UpdateImages

/*
*=================GET METHOD HANDLERS==================
*/

func GetAccounts(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})
	
	IDQ := c.Query("ID")
	idq := c.Query("id")
	
	//! ADD Q STRINGS FOR fname, lname, fullname, email, username
	//! ADD SPECIAL Q STRING FOR SORT
	if len(IDQ) > 0{
		query["ID"] = IDQ
	}else if len(idq) > 0{
		query["id"] = idq
	}

	c.IndentedJSON(http.StatusOK, db.Accounts_GET(query))
}

func GetPosts(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})

	IDQ := c.Query("ID")
	idq := c.Query("id")
	//! ADD Q STRINGS FOR fname, lname, fullname, email, username
	//! ADD SPECIAL Q STRING FOR SORT
	if len(IDQ) > 0{
		query["ID"] = IDQ
	}else if len(idq) > 0{
		query["id"] = idq
	}

	c.IndentedJSON(http.StatusOK, db.Posts_GET(query))
}

func GetPostsFullContext(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})
	//! ADD Q STRINGS FOR fname, lname, fullname, email, username
	//! ADD SPECIAL Q STRING FOR SORT
	IDQ := c.Query("ID")
	idq := c.Query("id")
	
	if len(IDQ) > 0{
		query["ID"] = IDQ
	}else if len(idq) > 0{
		query["id"] = idq
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

// func NewAccount(c *gin.Context){
// 	c.Header("Access-Control-Allow-Origin", "*")
	
// 	var newAcc account

// 	if err := c.BindJSON(&newAcc); err != nil{
// 		c.IndentedJSON(http.StatusBadRequest, nil)
// 		return
// 	}

// 	err := db.CreateNewAccount(newAcc)
// 	if err != nil{
// 		if _, ok := err.(*s.AccountExistsError); ok{
// 			c.IndentedJSON(http.StatusBadRequest, nil)
// 		}else{
// 			c.IndentedJSON(http.StatusInternalServerError, nil)
// 		}
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, newAcc)
// }



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
