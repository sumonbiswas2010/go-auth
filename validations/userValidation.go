package validations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func SignUp(c *gin.Context) {
	fmt.Println("Helllllllo from validations")
	var body struct {
		Email    string
		Password string
		Name     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid request",
		})
		return
	}

	if body.Password == "" {
		respondWithError(c, 401, "password required from validations. abort")
		return
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "password required from validations",
		})
		return
	}
	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Name required",
		})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Email required",
		})
		return
	}

	c.Next()
	return

}
