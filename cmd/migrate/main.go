package main

import (
	"fmt"

	"github.com/haisabdillah/golang-auth/core/infrastructure/db"
	"github.com/haisabdillah/golang-auth/core/models"
)

func main() {
	db, err := db.InitMySql()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	db.AutoMigrate(
		&models.Permission{},
		&models.Role{},
		&models.User{},
	)
}
