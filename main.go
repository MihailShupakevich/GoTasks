package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := "host=localhost user=admin dbname=GoDB password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})

	r := gin.Default()

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
	r.POST("/task", func(c *gin.Context) {
		var task Task
		c.BindJSON(task)
		if err != nil {
			return
		}
		if task.Text == "" {
			c.JSON(400, gin.H{"error": "Text is empty!"})
			return
		}
		tasks = append(tasks, task)
		db.Create(&task)
	})

	//обновление таски?
	r.PATCH("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		var task Task
		var updateTask Task
		if err := c.BindJSON(&updateTask); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		if updateTask.Text == "" {
			c.JSON(400, gin.H{"error": "Text is empty!"})
			return
		}
		if err := db.Model(&task).Where("id = ?", id).Update("checkbox", updateTask.Checkbox).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update task"})
			return
		}
		c.JSON(200, gin.H{"message": "Task updated successfully"})

	})

	//delete task
	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.First(&Task{}, id)
		db.Delete(&Task{}, result)
		if result.Error != nil {
			c.JSON(401, gin.H{"error": result.Error.Error()})
		}
	})

	r.Run(":8080")
}
