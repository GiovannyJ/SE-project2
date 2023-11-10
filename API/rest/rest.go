package rest


import (
	"github.com/gin-gonic/gin"
)

func API(){
	router := gin.Default()
	
	/*
	*=========================GET METHODS================================
	*/
	router.GET("/account", GetAccounts)
	router.GET("/posts", GetPosts)
	router.GET("/posts/fullcontext", GetPostsFullContext)

	/*
	*=========================POST METHODS================================
	*/
	router.POST("/account/new", NewAccount)
	router.POST("/posts/create", NewPost)
	router.POST("/login",Login)

	/*
	*=========================PATCH METHODS================================
	*/
	// router.PATCH("/account/update", UpdateAcc)
	// router.PATCH("/posts/update",UpdatePost)


	/*
	*=========================DELETE METHODS================================
	*/
	// router.DELETE("/account/delete:/THING", DelAccount)
	// router.DELETE("/posts/delete:/THING", DelPost)


	//*activate the server
	router.Run("localhost:8080")
}