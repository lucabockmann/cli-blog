package helper

import "github.com/gin-gonic/gin"

// | ----------------------------
// | Funktion: SetResponseHeader
// | ----------------------------
// | Override the Response Header for reuse in the routes.go
// | file and set it for each route
// | ----------------------------
func SetResponseHeader(origin string, method string, c *gin.Context) {
	// | ----------------------------
	// | Use the first parameter to set the allowed origin
	// | which can send requests to out api route
	// |
	// | allowed: ["*", "<origin>, null"]
	// | ----------------------------
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

	// | ----------------------------
	// | Use the second parameter to set the Allowed Method in Response Header.
	// | This can be set for each route seperate
	// |
	// | allowed: ["GET", "POST", "DELETE", "HEAD", "OPTIONS", "PUT"]
	// | ----------------------------
	c.Writer.Header().Set("Access-Control-Allow-Methods", method)
}
