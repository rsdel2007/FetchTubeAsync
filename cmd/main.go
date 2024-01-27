package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rsdel2007/proj/internal/api"
	"github.com/rsdel2007/proj/internal/job"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initliatise sqlite db
	godotenv.Load()
	db, err := gorm.Open(sqlite.Open("youtube.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Create JobQueue with DatabaseInitializer
	jobQueue := job.NewJobQueue(db)

	go jobQueue.StartFetchingJob()

	go func() {
		router := mux.NewRouter()
		api.RegisterHandlers(db, router)

		port := ":8081"
		log.Printf("Server started on port %s", port)
		log.Fatal(http.ListenAndServe(port, router))
	}()

	select {}
}
