package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:<senha>@/<nome bd>")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Zoldyck")
	stmt.Exec(2001, "Freecs")
	aux, err := stmt.Exec(1, "Illumi") //chave duplicada

	fmt.Println(aux)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
