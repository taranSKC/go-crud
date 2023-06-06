package controllers

import (
	initializers "crud/intializers"
	"crud/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreateController(c *gin.Context) {
	var Body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&Body)

	post := models.Post{Title: Body.Title, Body: Body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured in the system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": post,
	})

	fmt.Println("Printing to console")
}

func GetPostsController(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK,
		gin.H{
			"posts": posts,
		})

}

func GetPostController(c *gin.Context) {
	id := c.Param("post-id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func UpdatePostController(c *gin.Context) {
	var Body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&Body)
	var post models.Post
	id := c.Param("post-id")
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: Body.Title, Body: Body.Body})
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePostController(c *gin.Context) {

	id := c.Param("post-id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}
