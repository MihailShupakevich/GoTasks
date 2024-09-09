package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source/postgres"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "пользователь:пароль@localhost/база_данных")
	if err != nil {
		log.Fatal(err)
	}

	// Создание источника миграции
	sourceInstance, err := postgres.NewDatabaseInstance(
		"postgres",
		&postgres.Config{
			DatabaseName: "база_данных",
			User:         "пользователь",
			Password:     "пароль",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Создание миграции
	m, err := migrate.NewWithInstance("file:///migrations", sourceInstance)
	if err != nil {
		log.Fatal(err)
	}

	// Применение миграции
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Миграция применена")
}
