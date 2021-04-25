package database

import (
	"../models"
	"database/sql"
	"fmt"
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

var DB2 *sql.DB

func Connect2() {
	db, err := sql.Open("godror", "gomez/gomez@localhost:1521/ORCLCDB.localdomain")
	if err != nil {
		fmt.Println(err) //Could not connect to the database
		return
	}
	DB2 = db
	//defer db.Close()
	fmt.Println("Connect to the database successful")
}

func ExecuteQuery(query string) (*sql.Rows, error) {
	rows, err := DB2.Query(query)
	return rows, err
}