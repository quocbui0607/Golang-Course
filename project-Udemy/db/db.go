package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseInstance struct {
	db *sql.DB
}

var databaseInstance DatabaseInstance

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:lucario156@tcp(127.0.0.1:3306)/golang_project?parseTime=true")

	if err != nil {
		panic("Could not connect to database.")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(256) NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create user table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = db.Exec(createEventsTable)

	if err != nil {
		panic("Could not create event table.")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			event_id INTEGER,
			user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = db.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}

func GetInstance() *sql.DB {
	return databaseInstance.db
}

func SetDatabaseInstance(database *sql.DB) DatabaseInstance {

	if databaseInstance.db == nil {
		databaseInstance = DatabaseInstance{
			db: database,
		}
	}

	return databaseInstance
}
