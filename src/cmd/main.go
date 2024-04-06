package main

import (
	"log"
	"net/http"
)

func init() {
	db.InitDB()
}

func main() {
	//http.HandleFunc("/post-data", handlePost)

	// Start HTTP server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -X POST http://localhost:8080/albums/add \
// -H "Content-Type: application/json" \
// -d '{
//     "city": "Jaipur",
//     "temp": 35.5,
//     "humidity": 60,
//     "timestamp": "2024-04-06T14:30:00Z"
// }'
