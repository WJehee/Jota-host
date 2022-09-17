package controllers

import(
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/wjehee/jota-host/models"
)

// Register a new team
func Register(c *gin.Context) {
    team := models.Team{
        Username: "bob",
        Password: "test",
    }
    models.DB.Create(&team)
    c.JSON(http.StatusCreated, team)
}

// Login
func Login(c *gin.Context) {

}

func GetTeam(c *gin.Context) {
    var team models.Team
    models.DB.First(&team, "username = ?", "bob")
    c.JSON(http.StatusOK, team)
}

