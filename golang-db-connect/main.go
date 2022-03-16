package main

import (
	"database/sql"
	"fmt"
	//blank identifier
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=556690 dbname=test_db")

	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	fmt.Println("connected!")

	//create a table-- not working why?
	// _, err = db.Exec(`
	//   CREATE TABLE IF NOT EXISTS accounts(
	// 		account_id SERIAL PRIMARY KEY,
	// 		user_name TEXT,
	// 		user_id SERIAL NOT NULL
	// 	);
	// `)

	//insert some data
	// account_id := 1
	// user_name := "yang lyu"
	// user_id := 1
	// _, err = db.Exec(`
	//   INSERT INTO accounts (account_id, user_name, user_id)
	// 	VALUES ($1, $2);`, account_id, user_name, user_id)

	//insert some data
	// email := "abc@124.com"
	// name := "yang lyu"
	// id := 2
	// _, err = db.Exec(`
	//   INSERT INTO customers (name, email, id)
	// 	VALUES ($1, $2, $3);`, name, email, id)

	// checkErr(err)

	//query a data
	_, err = db.Exec(`
	  SELECT * FROM customers WHERE id=1;`)

	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
