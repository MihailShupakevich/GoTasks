package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	ID    int
	Email string
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=GoDB user=admin password=admin port=5432")
	if err != nil {
		return nil, err
	}
	return db, nil
}

var users []User

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/users", func(c *gin.Context) {

		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to query database"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Email)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to scan row"})
				return
			}
			users = append(users, user)
		}

		c.JSON(200, gin.H{"users": users})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": "Неверный запрос"})
			return
		}

		// Check if the user struct is empty
		if user.Email == "" {
			c.JSON(400, gin.H{"error": "Электронная почта обязательна"})
			return
		}

		_, err = db.Exec("INSERT INTO users (email) VALUES ($1) RETURNING id", user.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Не удалось выполнить запрос к базе данных"})
			return
		}
		users = append(users, user)
		c.JSON(201, gin.H{
			"message": "Successfully create user",
		})
	})

	r.Run(":8080")

}
