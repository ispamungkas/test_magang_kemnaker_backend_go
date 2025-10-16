package utils

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func waitForDB(dsn string, retries int) {
	for i := 0; i < retries; i++ {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			if err := db.Ping(); err == nil {
				db.Close()
				return
			}
		}
		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Database not ready after retries")
}

func InitDB(url string) {
	// waitForDB(url, 10)

	m, err := migrate.New(
		"file://./migrations",
		url,
	)
	if err != nil {
		log.Println("Failed to create migrate instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("Migration failed:", err)
	}

	log.Println("Database migration completed!")
}
