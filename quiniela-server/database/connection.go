package database

import (
	"../models"
	"github.com/cengsin/oracle"
	_ "github.com/godror/godror"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(oracle.Open("gomez/gomez@localhost:1521/ORCLCDB.localdomain"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	connection.AutoMigrate(&models.ADMIN{})
	DB = connection
}

/*
	db, err := sql.Open("godror", "gomez/gomez@localhost:1521/ORCLCDB.localdomain")
	if err != nil {
		fmt.Println(err) //Could not connect to the database
		return
	}
	defer db.Close()

	fmt.Println("Connect to the database successful")

*/
