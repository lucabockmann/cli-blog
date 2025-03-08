package helper

import "github.com/gin-gonic/gin"

func SetRequestHeader(origin string, method string, c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Methods", method)
}
