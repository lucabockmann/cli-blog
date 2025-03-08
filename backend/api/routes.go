package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucabockmann/go_boilerplate/helper"
	"github.com/lucabockmann/go_boilerplate/resources"
)

func CreateRoutes() {
	r := gin.Default()

	r.GET("/api/v1/get/posts", getPosts)

	r.Run(":8080")
}

func getPosts(c *gin.Context) {
	posts := []resources.PostStruct{}

	helper.SetRequestHeader("*", "GET", c)

	err := db.Select(&posts, "SELECT * FROM table")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, posts)
}
