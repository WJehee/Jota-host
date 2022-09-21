package models

type Team struct {
    Username string `json:"username" binding:"min=2,max=32,required" gorm:"primaryKey"`
    Password string `json:"password" binding:"min=8,max=128,required"`
    Points int64 `gorm:"default:0"`
}
