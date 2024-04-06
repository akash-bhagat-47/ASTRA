package main

import (
	db "ASTRA/src/database"
	"ASTRA/src/handlers"
	"log"

	"github.com/labstack/echo"
)

func init() {
	db.InitDB()
}

func main() {
	e := echo.New()
	e.POST("/ingest", handlers.IngestData)

	log.Println("Server listening on :8080")
	log.Fatal(e.Start(":8080"))
}
