package watcher

import (
	db "ASTRA/src/database"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func MonitorFiles() {
	directory := "/tmp/astra"
	log.Println("Initiated monitoring on directory", directory)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from PANIC!!! Error was:", err)
		}
	}()

	// Create a new watcher instance
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
	}
	defer watcher.Close()

	err = watcher.Add(directory)
	if err != nil {
		log.Println(err)
		return
	}

	// Start processing events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// Check if the event is a create event for a new file
			if event.Op&fsnotify.Create == fsnotify.Create {
				log.Println("New file created:", event.Name)
				updateDbAndDelete(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func updateDbAndDelete(filename string) {
	// Read the contents of the file
	rawData, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	var payload []db.DataLine

	if err = json.Unmarshal(rawData, &payload); err != nil {
		return
	}

	if err = updateDB(payload); err != nil {
		return
	}

	// Delete the file
	err = os.Remove(filename)
	if err != nil {
		log.Println("Error deleting file:", err)
		return
	}

	log.Println("File deleted:", filename)
}

func updateDB(records []db.DataLine) error {

	var wg sync.WaitGroup
	numWorkers := 5
	recordsChan := make(chan db.DataLine)

	//spawning workers here
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go updateWorker(recordsChan, &wg)
	}

	//sending actual fields of payload to the channel
	for _, record := range records {
		recordsChan <- record
	}
	close(recordsChan)
	wg.Wait()
	return nil
}

func updateWorker(records <-chan db.DataLine, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range records {
		if err := db.Insert(record); err != nil {
			log.Println(err)
			return
		}
	}
}
