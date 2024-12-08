package main

import (
	"awesomeProject/cmd/docs"
	"awesomeProject/internal/api/handlers"
	"awesomeProject/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/swaggo/files" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title ToDo API
// @version 1.0
// @description Description of the ToDO REST API
// @Summary      List tasks
// @Description  get tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Router       /tasks [get]
// @host localhost:8080
// @BasePath /tasks
func main() {
	db1, err := db.Сonnection()
	if err != nil {
		log.Fatal(err)
	}

	serve := gin.New()

	routes := serve.Group("/tasks")
	if db1 != nil {
		taskHandler := handlers.TaskHandler{DB: db1}
		routes.GET("/", taskHandler.FindAll)
		routes.POST("/", taskHandler.Post)
		routes.PATCH("/:id", taskHandler.Update)
		routes.DELETE("/:id", taskHandler.Delete)
		routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} else {
		log.Fatal("Failed to establish database connection")
	}

	docs.Init()

	serve.Run(":8080")
}
