package main

import (
	"fmt"
	"go-auth/controllers"
	"go-auth/initializers"
	"go-auth/validations"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	// initializers.SyncDB()
}
func main() {
	fmt.Println("Hello, world 2!")
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.POST("/signup", validations.SignUp, controllers.SignUp)
	r.POST("/login", validations.Login, controllers.Login)
	r.GET("/login", validations.CheckToken, controllers.CheckLogin)
	r.Run()
}
