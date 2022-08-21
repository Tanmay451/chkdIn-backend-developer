package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenAuth redirect to login page if user is not login
func TokenAuth(c *gin.Context) {
	tokenString, err := c.Cookie("tokenString")
	if err != nil {
		log.Println("TokenAuth: failed while getting token string from cookie with error: ", err)
		AuthFailed := "/api/auth-failed"
		c.Redirect(http.StatusFound, AuthFailed)
		return
	}
	fmt.Println("tokenString: ", tokenString)
	c.Next()
}
