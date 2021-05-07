package database

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

/*var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(oracle.Open("gomez/gomez@localhost:1521/ORCLCDB.localdomain"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	connection.AutoMigrate(&models.ADMIN{})
	DB = connection
}*/

var DB2 *sql.DB

func Connect() {
	db, err := sql.Open("godror", "gomez/gomez@localhost:1521/ORCLCDB.localdomain")
	if err != nil {
		fmt.Println(err) //Could not connect to the database
		return
	}
	DB2 = db
	fmt.Println("Connect to the database successful")
	//defer db.Close()
}

func ExecuteQuery(query string) (*sql.Rows, error) {
	rows, err := DB2.Query(query)
	//defer rows.Close()
	fmt.Println(err)
	return rows, err
}

func Execute(query string)  error {
	_,err := DB2.Exec(query)
	fmt.Println("EXECUTE: ", err)
	return err
}