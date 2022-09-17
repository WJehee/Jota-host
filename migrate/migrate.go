package main

import "github.com/wjehee/jota-host/models"

func main() {
    models.DB.AutoMigrate(&models.Team{})
}

