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
	router.GET("/posts/:id/comments", GetPostsComments)

	/*
	*=========================POST METHODS================================
	*/
	router.POST("/account/new", NewAccount)
	router.POST("/posts/create", NewPost)
	router.POST("/posts/:id/comment", NewComment)
	router.POST("/login",Login)

	/*
	*=========================PATCH METHODS================================
	*/
	router.PATCH("/account/update", UpdateAcc)
	router.PATCH("/posts/update",UpdatePost)
	router.PATCH("/posts/:id/comment", UpdateComment)

	/*
	*=========================DELETE METHODS================================
	*/
	router.DELETE("/account/delete/:id", DelAccount)
	router.DELETE("/posts/delete/:id", DelPost)
	router.DELETE("/posts/:id/delete/:id", DelComment)


	//*activate the server
	router.Run("localhost:8080")
}