package controllers

import (
	"encoding/json"
	"fmt"
	"go-auth/initializers"
	"go-auth/models"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

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
	fmt.Println("Helllllllo")
	c.JSON(http.StatusAccepted, gin.H{
		"status": "true",
		"error":  "create user done",
	})
	return

}
