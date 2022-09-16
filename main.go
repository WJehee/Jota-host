package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wjehee/jota-host/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
      log.Fatal("failed to connect to database")
    }
    db.AutoMigrate(&models.Team{})
}


func main() {
    server := gin.New()
    server.Use(gin.Recovery())
    server.Use(gin.Logger())

    server.GET("/", home)
    server.GET("/register", register)
    server.GET("/team", getTeam)
    server.POST("/submit", submit)

    server.Run()
}

func home(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "hello world!",
    })
}

// Register a new team
func register(c *gin.Context) {
    team := models.Team{
        Username: "bob",
        Password: "test",
    }
    db.Create(&team)
    c.JSON(http.StatusCreated, team)
}

func getTeam(c *gin.Context) {
    var team models.Team
    db.First(&team, "username = ?", "bob")
    c.JSON(http.StatusOK, team)
}

// Submit a question
func submit(c *gin.Context) {

}

