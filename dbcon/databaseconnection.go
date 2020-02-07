package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Create(db *sql.DB, fName, lName string) (int64, error)  {
	insertString := "insert into  Person (LastName, FirstName) values ('" + lName + "', '" + fName + "');"

	result, err := db.Exec(insertString)

	if err != nil {
		fmt.Println("Insert failed")
		return -1, err
	}
	// db.Close()

	return result.LastInsertId();
}

func ReadWithFirstName(db *sql.DB, fName string) (int, error)  {
	queryStirng := "SELECT * FROM Person where FirstName = '" + fName + "';"
	rows, err := db.Query(queryStirng)

	if err != nil {
		fmt.Println("Read failed")
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var fName, lName string
		var personId int

		err := rows.Scan(&personId, &lName, &fName)

		if err != nil {
			fmt.Println("Row Read Failed")
			return -1, err
		}
		fmt.Printf("First Name: %s, Last Name: %s\n", fName, lName)
		count++
	}

	return count, nil

}

func main() {
	//conString := "server=10.5.205.54;user id=root;password=root;port=3306;database=testdata_gbhat;"
	conString := "root:root@tcp(10.5.205.54:3306)/testdata_gbhat"

	conn, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return
	}
	fmt.Println("Connected!")
	defer conn.Close()

	// Create Person
	/*createID, err := Create(conn, "Jake", "United States")
	if err != nil {
		log.Fatal("Create failed:", err.Error())
	}
	fmt.Printf("Inserted ID: %d successfully.\n", createID)*/

	// Read Person
	count, err := ReadWithFirstName(conn, "Jake")
	if err != nil {
		log.Fatal("ReadEmployees failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", count)

}
