package rest


import (
	"github.com/gin-gonic/gin"
)

func API(){
	router := gin.Default()
	
	/*
	*=========================GET METHODS================================
	*/
	router.GET("/account", GetAccounts) //W query strings
	router.GET("/posts", GetPosts) //W query strings
	router.GET("/posts/fullcontext", GetPostsFullContext)

	/*
	*=========================POST METHODS================================
	*/
	// router.POST("/account/new", NewAccount)
	// router.POST("/login",Login)
	// router.POST("/login/guest", LoginGuest)
	// router.POST("/posts/create", NewPost)

	/*
	*=========================PATCH METHODS================================
	*/
	// router.PATCH("/account/update", UpdateAcc)
	// router.PATCH("/posts/update",UpdatePost)


	/*
	*=========================DELETE METHODS================================
	*/
	// router.DELETE("/account/delete", DelAccount)
	// router.DELETE("/posts/delete", DelPost)


	//*activate the server
	router.Run("localhost:8080")
}