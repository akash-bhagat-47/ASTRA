package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func InitDB() {
	fmt.Println("Initialising DB connection")
	dsn := "host=localhost user=akash.b password= dbname=weatherStation port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("dberror", err)
		panic(err)
	}

	if err := dbConn.AutoMigrate(&DataLine{}); err != nil {
		fmt.Println("AutoMigrate Error", err)
		panic(err)
	}
	fmt.Println("Done with automigration!!")
	DbConn = dbConn
}

func ConnectToDB() *gorm.DB {
	return DbConn
}
