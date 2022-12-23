package controllers

import (
	"encoding/json"
	"fmt"
	"go-auth/initializers"
	"go-auth/models"
	"go-auth/validations"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	body := c.MustGet("body").(validations.SignupData)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "password hash error",
		})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	res := initializers.DB.Create(&user)

	// spew.Dump(res)
	if res.Error != nil {
		fmt.Println("dump")
		spew.Dump(res.Error)
		// var pgErr *pgconn.PgError
		// var errMsg = ""
		// if errors.As(err, &pgErr) {
		// 	fmt.Println("hello from inside")
		// 	spew.Dump(pgErr)
		// 	// fmt.Println(pgErr.Message)
		// 	// fmt.Println(pgErr.Code)
		// 	errMsg = string(pgErr.Message)
		// }
		b, _ := json.Marshal(res.Error)
		m := make(map[string]interface{})
		err := json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m["Code"])
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  m["Message"],
		})
		return
	}
	// var user models.User
	// initializers.DB.First(&user, "email=?", body.Email)
	// spew.Dump(user)

	// if user.ID == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": "failed",
	// 		"error":  "invalid email",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusAccepted, gin.H{
		"status": "true",
		"error":  "create user done",
	})
	return

}
func Login(c *gin.Context) {

	body := c.MustGet("body").(validations.LoginData)

	var user models.User
	initializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid pass",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"name": user.Name,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		spew.Dump(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid token generation",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "done",
		"token":  tokenString,
	})
	return
}

func CheckLogin(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	c.JSON(http.StatusAccepted, gin.H{
		"status": "done",
		"token":  "Logged In",
		"user": userData["name"],
	})

}
