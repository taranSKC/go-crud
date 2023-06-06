package main

import (
	controllers "crud/controllers"
	initializers "crud/intializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/post", controllers.PostsCreateController)
	r.GET("/posts", controllers.GetPostsController)
	r.GET("/posts/:post-id", controllers.GetPostController)
	r.POST("/posts/:post-id", controllers.UpdatePostController)
	r.DELETE("/posts/:post-id", controllers.DeletePostController)
	r.Run()

}
