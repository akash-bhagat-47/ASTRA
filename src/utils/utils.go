package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func WriteToFile(data []byte) error {
	dir := "/tmp/astra"

	if err := dirExists(dir); err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.json", time.Now().UnixNano())
	filePath := filepath.Join(dir, filename)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Println("Error writing data to file")
		return err
	}

	fmt.Printf("Data saved to %s", filePath)
	fmt.Printf("\n")
	return nil
}

func dirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Directory does not exist, create it
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Println("Failed to create directory:", err)
			return err
		}
		log.Println("Directory created successfully:", dir)
	} else if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}
