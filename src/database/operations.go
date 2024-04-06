package db

import "log"

func Insert(data DataLine) error {
	DbConnection := ConnectToDB()
	if err := DbConnection.Create(&data).Error; err != nil {
		log.Println("Failed to insert data:", err)
		return err
	}

	log.Println("Record inserted successfully in db !!!")
	return nil
}
