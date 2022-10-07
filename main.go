package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/wjehee/jota-host/controllers"
	"github.com/wjehee/jota-host/middleware"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Could not load .env")
    }
    secret := os.Getenv("SECRET")
    server := gin.New()
    server.Use(gin.Recovery())
    server.Use(gin.Logger())
    store := cookie.NewStore([]byte(secret))
    server.Use(sessions.Sessions("session", store))

    // Account routes
    server.POST("/register", controllers.Register)
    server.POST("/login", controllers.Login)
    apiRoutes := server.Group("/api", middleware.Authenticate())
    {
        apiRoutes.POST("/submit", submit)
    }
    server.GET("/team", controllers.GetTeams)
    server.GET("/team/:username", controllers.GetTeam)

    server.Run()
}

// Submit a question
func submit(c *gin.Context) {

}

