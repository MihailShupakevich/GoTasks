package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"type:text"`
	Checkbox  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

var tasks []Task

func main() {
	dsn := "host=localhost user=admin dbname=GoDB password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Task{})

	//получение всех записей
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		db.Find(&tasks)
		json.NewEncoder(w).Encode(tasks)
	})

	//создание 1 записи
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := db.Create(&task)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		if task.Text == "" {
			http.Error(w, "Empty task text", http.StatusBadRequest)
		}
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode("task succesfully created")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//var users []User
//
//func main() {
//	db, err := ConnectDB()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	r := gin.Default()
//
//	r.GET("/", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "Hello, World!",
//		})
//	})
//
//	r.GET("/users", func(c *gin.Context) {
//
//		rows, err := db.Query("SELECT * FROM users")
//		if err != nil {
//			c.JSON(500, gin.H{"error": "Failed to query database"})
//			return
//		}
//		defer rows.Close()
//
//		for rows.Next() {
//			var user User
//			err := rows.Scan(&user.ID, &user.Email)
//			if err != nil {
//				c.JSON(500, gin.H{"error": "Failed to scan row"})
//				return
//			}
//			users = append(users, user)
//		}
//
//		c.JSON(200, gin.H{"users": users})
//	})
//
//	r.POST("/user", func(c *gin.Context) {
//		var user User
//		err := c.BindJSON(&user)
//		if err != nil {
//			c.JSON(400, gin.H{"error": "Неверный запрос"})
//			return
//		}
//
//		// Check if the user struct is empty
//		if user.Email == "" {
//			c.JSON(400, gin.H{"error": "Электронная почта обязательна"})
//			return
//		}
//
//		_, err = db.Exec("INSERT INTO users (email) VALUES ($1) RETURNING id", user.Email)
//		if err != nil {
//			c.JSON(500, gin.H{"error": "Не удалось выполнить запрос к базе данных"})
//			return
//		}
//		users = append(users, user)
//		c.JSON(201, gin.H{
//			"message": "Successfully create user",
//		})
//	})
//
//	r.Run(":8080")
//
//}
