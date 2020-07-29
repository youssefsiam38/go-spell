package middlewares

import (
	"net/http"
	"github.com/youssefsiam38/spell/utils"
	"strings"
	"github.com/gin-gonic/gin"
)

// Auth to make authentication required
func Auth(c *gin.Context) {
		// Get Authorization token
		auth := c.GetHeader("Authorization")
		auth = strings.Replace(auth, "Bearer ", "", 1)
		
		retrievedUser, err := utils.VerifyJWT(auth)
	
		if err != nil {
			utils.ErrResponse(c, http.StatusUnauthorized, "Something went wrong with authorization please login")
			c.Abort()
		} else {
			c.Set("userPtr", retrievedUser)
			c.Next()
		}
	
}