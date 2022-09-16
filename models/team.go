package models

type Team struct {
    Username string `json:"username" binding:"min=2,max=32,required"`
    Password string `json:"password" binding:"min=8,max=128,required"`
    // Teamname string `json:"teamname" binding:"min=2, max=32"`
    // Points int64
}
