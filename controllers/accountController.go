package controllers

import (
	"database/sql"
	"fakhry/go-sql/entities"
	"fmt"
	"log"
)

func GetAllAccountsController(db *sql.DB) {
	// READ / SELECT data
	// var untuk menyimpan data yang dibaca di query select
	var accounts []entities.Account

	// menjalankan perintah query select
	rows, errSelect := db.Query("SELECT id, username, name, bio, join_date, location, email, password, created_at FROM accounts")
	// handle ketika terjadi error saat menjalankan select
	if errSelect != nil {
		log.Fatal("error run query select ", errSelect.Error())
	}

	//proses membaca per baris
	// menggunakan for karena data yang akan dibaca itu banyak
	for rows.Next() {
		var dataAccount entities.Account // var untuk menyimpan data akun per baris
		// var joinDate string
		// proses scan dan mapping data ke var dataAccount
		errScan := rows.Scan(&dataAccount.ID, &dataAccount.Username, &dataAccount.Name, &dataAccount.Bio, &dataAccount.JoinDate, &dataAccount.Location, &dataAccount.Email, &dataAccount.Password, &dataAccount.CreatedAt)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		// fmt.Println("joinDate:", joinDate)
		// memasukkan dataAccount ke accounts
		accounts = append(accounts, dataAccount)
	}

	for _, v := range accounts {
		fmt.Printf("username: %v, join-date: %v\n", v.Username, v.JoinDate)
	}
}
