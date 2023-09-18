package main

import (
	"dog-app-api/db"
	"dog-app-api/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Migrate成功")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Dog{})
}
