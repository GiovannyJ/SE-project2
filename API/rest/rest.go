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
	router.GET("/account", GetAccounts)

	/*
	*=========================POST METHODS================================
	*/
	router.POST("/account/new", NewAccount)

	/*
	*=========================PATCH METHODS================================
	*/
	router.PATCH("/account/update", UpdateAcc)

	//*activate the server
	router.Run("localhost:8080")
}