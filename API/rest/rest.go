package rest


import (
	"github.com/gin-gonic/gin"
)

func API(){
	router := gin.Default()
	
	/*
	*=========================GET METHODS================================
	*/
	// router.GET("/get/data", GET_data)
	router.GET("/account/all", GetAccounts)

	/*
	*=========================POST METHODS================================
	*/
	router.POST("/post/data", POST_data)

	/*
	*=========================PATCH METHODS================================
	*/
	router.PATCH("/patch/data", PATCH_data)

	//*activate the server
	router.Run("localhost:8080")
}