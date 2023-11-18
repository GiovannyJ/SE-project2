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
	router.GET("/posts/:id/commentsfullcontext", GetPostsCommentsFullContext)

	/*
	*=========================POST METHODS================================
	*/
	router.POST("/account/new", NewAccount)
	router.POST("/posts/create", NewPost)
	router.POST("/comment/new", NewComment)
	router.POST("/login",Login)

	/*
	*=========================PATCH METHODS================================
	*/
	router.PATCH("/account/update", UpdateAcc)
	router.PATCH("/posts/update",UpdatePost)
	router.PATCH("/comment/update", UpdateComment)

	/*
	*=========================DELETE METHODS================================
	*/
	router.DELETE("/account/delete/:id", DelAccount)
	router.DELETE("/posts/delete/:id", DelPost)
	router.DELETE("/posts/delete/comment/:id", DelComment)

	/*
	*=========================FILE METHODS================================
	*/
	router.GET("/uploads/:file", RetrieveFile)
	router.GET("/fileinfo", GetFileInfo)
	router.POST("/uploads", UploadFile)
	router.PATCH("/filein")


	//*activate the server
	router.Run("localhost:8080")
}