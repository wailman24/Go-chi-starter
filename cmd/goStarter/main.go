package main

import (
	"fmt"
	"net/http"

	"github.com/wailman24/Go-chi-starter.git/internal/models"
	"github.com/wailman24/Go-chi-starter.git/internal/router"
	"github.com/wailman24/Go-chi-starter.git/pkg/db"
)

func main() {
	db.Connect()
	fmt.Println("connected ... ")
	err := db.Db.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println("Error migrating database:", err)
	} else {
		fmt.Println("Migration completed successfully...")
	}

	http.ListenAndServe(":8080", router.MainRoutes())
}
