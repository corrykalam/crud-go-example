package middlewares

import (
	"fmt"
	"net/http"
	"sesi12-final/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization"})
			c.Abort()
			return
		}

		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization"})
			c.Abort()
			return
		}
		if authSplit[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization"})
			c.Abort()
			return
		}

		token, err := helpers.VerifyToken(authSplit[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad koxak"})
			c.Abort()
			return
		}
		fmt.Println(token)

		c.Next()
	}
}
