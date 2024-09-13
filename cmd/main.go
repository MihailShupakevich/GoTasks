package main

import (
	"awesomeProject/internal/api/handlers"
	"awesomeProject/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db1, err := db.Сonnection()
	if err != nil {
		log.Fatal(err) // handle error when db connection fails
	}

	// Инициализация роутинга
	serve := gin.New()
	routes := serve.Group("/tasks")
	if db1 != nil {
		taskHandler := handlers.TaskHandler{DB: db1}
		routes.GET("/", taskHandler.FindAll)
		routes.POST("/", taskHandler.Post)
		routes.PATCH("/:id", taskHandler.Update)
		routes.DELETE("/:id", taskHandler.Delete)
	} else {
		log.Fatal("Failed to establish database connection")
	}
	serve.Run(":8080")
}
