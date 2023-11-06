package rest

import (
	db "API/conn"
	// s "API/models"
	// "encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)


/*
*=================GET METHOD HANDLERS==================
*/

func GetAccounts(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	var query = make(map[string]interface{})
	
	IDQ := c.Query("ID")
	
	if len(IDQ) > 0{
		query["ID"] = IDQ

	}
	c.IndentedJSON(http.StatusOK, db.Accounts_GET(query))
}

/*
*=================POST METHOD HANDLERS==================
*/
func POST_data(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
}

func samplePost(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	// var newAccount loginDATA

	// if err := c.BindJSON(&newAccount); err != nil {
	// 	//err handling
	// 	//c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	c.IndentedJSON(http.StatusBadRequest, nil)
	// 	return
	// }

	// if err := db.NewAccount(newAccount); err != nil {
	// 	//err handing
	// 	// c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	c.IndentedJSON(http.StatusInternalServerError, nil)
	// }
	
	// d := jsondata{Data: newAccount}

	// c.IndentedJSON(http.StatusCreated, d)
}

/*
*=================PATCH METHOD HANDLERS==================
*/
func PATCH_data(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
}

func samplePatch(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	// var upACC updateACC

	// if err := c.BindJSON(&upACC); err != nil {
	// 	//err handle
	// 	// c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	c.IndentedJSON(http.StatusBadRequest, nil)
	// }
	// if err := db.UpdateData(upACC.Old, upACC.New); err != nil {
	// 	//err handle
	// 	// c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	c.IndentedJSON(http.StatusBadRequest, nil)
	// }

	// d := jsondata{Data: updateACC{
	// 	New: upACC.New,
	// 	Old: upACC.Old,
	// },
	// }
	// c.IndentedJSON(http.StatusOK, d)
}