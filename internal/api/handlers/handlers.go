package handlers

import (
	"awesomeProject/internal/domain"
	_ "awesomeProject/internal/repository"
	interfaces "awesomeProject/internal/usecase/task"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	_ "gorm.io/gorm"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	taskUseCase interfaces.TaskUseCase
}

type Response struct {
	ID       uint   `copier:"must"`
	Text     string `copier:"must"`
	Checkbox bool   `copier:"must"`
}

func (cr *TaskHandler) FindAll(c *gin.Context) {
	tasks, err := cr.taskUseCase.FindAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &tasks)

		c.JSON(http.StatusOK, response)
	}
}

//update task

func (cr *TaskHandler) Update(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	var updateTask domain.Task
	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := cr.taskUseCase.FindByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	task.Text = updateTask.Text
	task.Checkbox = updateTask.Checkbox
	if err := cr.taskUseCase.Update(c.Request.Context(), task); err != nil {
	}

	c.JSON(200, gin.H{"message": "Task updated successfully"})
}

// post task
func (cr *TaskHandler) Post(c *gin.Context) {
	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	task, err := cr.taskUseCase.Post(c.Request.Context(), newTask)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &task)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *TaskHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	task, err := cr.taskUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	cr.taskUseCase.Delete(ctx, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task is deleted successfully"})
}

////func handlers() {
////	r := gin.Default()
////	r.Use(enableCORS)
////	r.GET("/", func(c *gin.Context) {
////		c.JSON(200, gin.H{
////			"message": "ChiVa<1",
////		})
////	})
//
//	//получение всех записей
//	r.GET("/tasks", func(c *gin.Context) {
//		db.Find(&tasks)
//		c.JSON(200, gin.H{"tasks": tasks})
//	})
//
//
//
//	//создание 1 записи
//	r.POST("/tasks", func(c *gin.Context) {
//		task := new(Task)
//		c.BindJSON(task)
//		if err != nil {
//			return
//		}
//		if task.Text == "" {
//			c.JSON(400, gin.H{"error": "Text is empty!"})
//			return
//		}
//		db.Create(&task)
//	})
//
//	//обновление таски?
//	r.PATCH("/tasks/:id", func(c *gin.Context) {
//		id := c.Param("id")
//		var task Task
//		result := db.First(&task, id)
//		if result.Error != nil {
//			c.JSON(404, gin.H{"error": "Task not found"})
//			return
//		}
//
//		var updateTask Task
//		if err := c.BindJSON(&updateTask); err != nil {
//			c.JSON(400, gin.H{"error": "Invalid request"})
//			return
//		}
//		if updateTask.Text == "" {
//			c.JSON(400, gin.H{"error": "Text is empty!"})
//			return
//		}
//		if err := db.Model(&task).Updates(map[string]interface{}{"Text": updateTask.Text, "Checkbox": updateTask.Checkbox}).Error; err != nil {
//			c.JSON(500, gin.H{"error": "Failed to update task"})
//			return
//		}
//		c.JSON(200, gin.H{"message": "Task updated successfully"})
//	})
//
//	//delete task
//	r.DELETE("/tasks/:id", func(c *gin.Context) {
//		id := c.Param("id")
//		result := db.First(&Task{}, id)
//		db.Delete(&Task{}, result)
//		if result.Error != nil {
//			c.JSON(401, gin.H{"error": result.Error.Error()})
//		}
//	})
//	r.Run(":8080")
//}
//func enableCORS(c *gin.Context) {
//	w := c.Writer
//	r := c.Request
//
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
//	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
//
//	if r.Method == "OPTIONS" {
//		w.WriteHeader(http.StatusNoContent)
//		return
//	}
//	c.Next()
//}
