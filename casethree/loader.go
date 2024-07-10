package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"casethree/helper"
)

type userData struct {
	Id             int
	Country        string
	CreditCardType string
	CreditCard     string
	FirstName      string
	LastName       string
}

type transactionData struct {
	Id       int
	IdUser   int
	TotalBuy int
}

func CreateTableUser(db *sql.DB) interface{} {
	stmt, err := db.Prepare("CREATE TABLE users (id INT PRIMARY KEY AUTO_INCREMENT, country VARCHAR(255) NOT NULL, credit_card VARCHAR(255), credit_card_type VARCHAR(255), first_name VARCHAR(255) NOT NULL, last_name VARCHAR(255) NOT NULL);")
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	helper.PanicIfError(err)

	return nil
}

func CreateTableTransactions(db *sql.DB) interface{} {
	stmt, err := db.Prepare("CREATE TABLE transactions (id INT PRIMARY KEY AUTO_INCREMENT, id_user INT NOT NULL, total_buy INT NOT NULL);")
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	helper.PanicIfError(err)

	return nil
}

func DeleteTableUser(db *sql.DB) interface{} {
	stmt, err := db.Prepare("DROP TABLE users;")
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	helper.PanicIfError(err)

	return nil
}

func LoadTransactionDatabase(db *sql.DB, filename string) interface{} {
	CreateTableTransactions(db)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		if record[0] == "id" && record[1] == "id_user" && record[2] == "total_buy" {
			continue
		}

		var data transactionData
		id, err := strconv.Atoi(record[0])
		helper.PanicIfError(err)
		idUser, err := strconv.Atoi(record[1])
		helper.PanicIfError(err)
		totalBuy, err := strconv.Atoi(record[2])
		helper.PanicIfError(err)
		data.Id = id
		data.IdUser = idUser
		data.TotalBuy = totalBuy

		fmt.Printf("Record: %+v\n", data)

		stmt, err := db.Prepare("INSERT INTO transactions (id, id_user, total_buy) VALUES (?, ?, ?)")
		helper.PanicIfError(err)
		defer stmt.Close()

		_, err = stmt.Exec(data.Id, data.IdUser, data.TotalBuy)
		if err != nil {
			log.Println("Error inserting data:", err)
		}

		fmt.Println("Successfully inserted data:", data)
	}

	fmt.Println("Successfully read all data from CSV file!")

	return nil
}

func LoadCSVToDatabase(db *sql.DB, filename string) interface{} {
	CreateTableUser(db)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		if record[0] == "id" && record[1] == "country" && record[2] == "credit_card_type" && record[3] == "credit_card_number" && record[4] == "first_name" && record[5] == "last_name" {
			continue
		}

		var data userData
		id, err := strconv.Atoi(record[0])
		helper.PanicIfError(err)
		data.Id = id
		data.Country = record[1]
		data.CreditCardType = record[2]
		data.CreditCard = record[3]
		data.FirstName = record[4]
		data.LastName = record[5]

		fmt.Printf("Record: %+v\n", data)

		stmt, err := db.Prepare("INSERT INTO users (id, country, credit_card_type, credit_card, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?)")
		helper.PanicIfError(err)
		defer stmt.Close()

		_, err = stmt.Exec(data.Id, data.Country, data.CreditCardType, data.CreditCard, data.FirstName, data.LastName)
		if err != nil {
			log.Println("Error inserting data:", err)
		}

		fmt.Println("Successfully inserted data:", data)
	}

	fmt.Println("Successfully read all data from CSV file!")

	return nil
}
