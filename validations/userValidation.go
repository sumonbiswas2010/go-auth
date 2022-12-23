package validations

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

type LoginData struct {
	Email    string
	Password string
}
type SignupData struct {
	Email    string
	Password string
	Name     string
}

func SignUp(c *gin.Context) {
	var body SignupData

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
	}
	if body.Name == "" {
		respondWithError(c, 401, "Name required from validations. abort")

		return
	}
	if body.Email == "" {
		respondWithError(c, 401, "Email required from validations. abort")

		return
	}
	c.Set("body", body)
	c.Next()
	return

}

func Login(c *gin.Context) {

	var body LoginData
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

	}
	if body.Email == "" {
		respondWithError(c, 401, "Email required from validations. abort")
		return
	}

	c.Set("body", body)

	c.Next()

}
func CheckToken(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 6 {
		respondWithError(c, 401, "Unauthorizedd")
		return
	}
	const BEARER_SCHEMA = "Bearer "
	tokenString := authHeader[len(BEARER_SCHEMA):]
	if len(tokenString) < 6 {
		respondWithError(c, 401, "Unauthorizedd")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("userData", claims)
		c.Next()
		return
	} else {
		fmt.Println(err)
	}
	respondWithError(c, 401, "Unauthorizedd")
	return

}
