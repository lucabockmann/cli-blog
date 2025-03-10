package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucabockmann/cliblog/helper"
	"github.com/lucabockmann/cliblog/resources"
)

func CreateRoutes() {
	r := gin.Default()

	r.GET("/api/v1/get/posts", getPosts)
	r.GET("/api/v1/get/post/:id", getPostById)

	r.GET("/api/v1/get/userData", getUserData)

	r.Run(":8080")
}

func getPosts(c *gin.Context) {
	posts := []resources.PostStruct{}

	helper.SetResponseHeader("http://localhost", "GET", c)

	err := db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, posts)
}

func getPostById(c *gin.Context) {
	post := []resources.PostStruct{}
	helper.SetResponseHeader("http://localhost", "GET", c)
	id := c.Param("id")

	err := db.Select(&post, "SELECT * FROM posts WHERE (PostID) IN (?)", id)
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
	}

	c.JSON(200, post)
}

func getUserData(c *gin.Context) {
	user := []resources.UserStruct{}
	helper.SetResponseHeader("http://localhost", "GET", c)

	err := db.Select(&user, "SELECT * FROM userInfo")
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
	}

	c.JSON(200, user)
}
