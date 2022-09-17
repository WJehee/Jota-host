package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wjehee/jota-host/controllers"
)


func main() {
    server := gin.New()
    server.Use(gin.Recovery())
    server.Use(gin.Logger())

    // Account routes
    server.POST("/register", controllers.Register)
    server.POST("/login", controllers.Login)
    server.GET("/team", controllers.GetTeam)

    // Question routes
    server.POST("/submit", submit)

    server.Run()
}

// Submit a question
func submit(c *gin.Context) {

}

