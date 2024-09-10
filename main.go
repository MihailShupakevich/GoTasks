package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"type:text" json:"text"`
	Checkbox  bool      `gorm:"type:boolean" json:"checkbox"`
	CreatedAt time.Time `gorm:"index default current_timestamp()"`
	UpdatedAt time.Time `gorm:"index default current_timestamp()"`
}

var tasks []Task

func main() {
	r := gin.Default()
	r.Use(enableCORS)
	//
	dsn := "host=localhost user=admin dbname=GoDB password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ChiVa<1",
		})
	})
	//получение всех записей
	r.GET("/tasks", func(c *gin.Context) {
		db.Find(&tasks)
		c.JSON(200, gin.H{"tasks": tasks})
	})

	//создание 1 записи
	r.POST("/tasks", func(c *gin.Context) {
		task := new(Task)
		c.BindJSON(task)
		if err != nil {
			return
		}
		if task.Text == "" {
			c.JSON(400, gin.H{"error": "Text is empty!"})
			return
		}
		db.Create(&task)
	})

	//обновление таски?
	r.PATCH("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		var task Task
		result := db.First(&task, id)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}

		var updateTask Task
		if err := c.BindJSON(&updateTask); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		if updateTask.Text == "" {
			c.JSON(400, gin.H{"error": "Text is empty!"})
			return
		}
		if err := db.Model(&task).Updates(map[string]interface{}{"Text": updateTask.Text, "Checkbox": updateTask.Checkbox}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update task"})
			return
		}
		c.JSON(200, gin.H{"message": "Task updated successfully"})
	})

	//delete task
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.First(&Task{}, id)
		db.Delete(&Task{}, result)
		if result.Error != nil {
			c.JSON(401, gin.H{"error": result.Error.Error()})
		}
	})

	r.Run(":8080")
}

func enableCORS(c *gin.Context) {
	w := c.Writer
	r := c.Request

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	c.Next()
}
