package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wjehee/jota-host/models"
	"golang.org/x/crypto/bcrypt"
)

// Register a new team
func Register(c *gin.Context) {
    var team models.Team;
    err := c.Bind(&team)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
    } else {
        // hash the password of the user before inserting into the database
        hashed_pw, err := bcrypt.GenerateFromPassword([]byte(team.Password), 10)
        if err != nil {
            c.Status(http.StatusInternalServerError)
        }
        team.Password = string(hashed_pw)
        result := models.DB.FirstOrCreate(&team)
        if result.RowsAffected == 0 {
            c.Status(http.StatusUnprocessableEntity)
        } else {
            c.Status(http.StatusCreated)
        }
    }
}

// Login
func Login(c *gin.Context) {
    var login struct {
        Username string
        Password string
    }
    err := c.Bind(&login)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
    }
    var team models.Team
    models.DB.First(&team, "username = ?", login.Username)
    err = bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(login.Password))
    if err != nil {
        c.Status(http.StatusForbidden)
    } else {
        session := sessions.Default(c)
        session.Options(sessions.Options{
            Secure: true,
        })
        session.Set("user", login.Username)
        session.Save()
        c.Status(http.StatusOK)
    }
}

func GetTeam(c *gin.Context) {
    username := c.Param("username")
    var team models.APITeam
    models.DB.Model(models.Team{}).First(&team, "username = ?", username)
    c.JSON(http.StatusOK, team)
}

func GetTeams(c *gin.Context) {
    var teams []models.APITeam
    models.DB.Model(&models.Team{}).Find(&teams)
    c.JSON(http.StatusOK, teams)
}


