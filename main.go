package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"KemnakerMagang/authentication"
	"KemnakerMagang/utils"

	_ "github.com/lib/pq"
)

func main() {

	// Initialize Database
	
	// Database configuration
	// 1.) user
	// 2.) password
	// 3.) host
	// 4.) port
	// 5.) database name

	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	log.Printf("postgres://%s:%s@%s:%v/%s?sslmode=disable\n",
		user, passwd, host, port, name)

	utils.InitDB("postgresql://postgres:fHrONrgVwzmWRGWDTNBwsRLaOOgZjknM@gondola.proxy.rlwy.net:46912/railway")

	db, dbErr := sql.Open("postgres", "postgresql://postgres:fHrONrgVwzmWRGWDTNBwsRLaOOgZjknM@gondola.proxy.rlwy.net:46912/railway")
	if dbErr != nil {
		log.Fatalf("Error connecting db: %v", dbErr)
	}

	// Tes koneksi
	if err := db.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	r := mux.NewRouter()

	// Listener server
	server := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Service
	authService := authentication.NewAuthenticationService(db)

	// Handler
	r.HandleFunc("/users", authService.GetAllUser).Methods("GET")
	r.HandleFunc("/users", authService.AddUser).Methods("POST")
	r.HandleFunc("/users/{id}", authService.GetUserById).Methods("GET")

	fmt.Printf("Listener server localhost:%v\n", server.Addr)
	server.ListenAndServe()
}
