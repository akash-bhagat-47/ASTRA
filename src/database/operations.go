package db

import "fmt"

func Insert(data DataLine) error {
	DbConnection := ConnectToDB()
	if err := DbConnection.Create(&data).Error; err != nil {
		fmt.Println("Failed to insert data:", err)
		return err
	}

	fmt.Println("Data inserted successfully!")
	return nil
}
