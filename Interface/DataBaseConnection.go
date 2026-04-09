package main

import "fmt"

type DatabaseConnection interface {
	connection()
}

type DataBaseConnectionImplementation struct {
}

func (e DataBaseConnectionImplementation) connection() {
	fmt.Println("database connection successfully")
}

func main() {

	connect := DataBaseConnectionImplementation{}
	var dbConnection DatabaseConnection = connect
	dbConnection.connection()
}
