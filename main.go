package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db        *sql.DB
	separator = " | "
)

func main() {
	fileName := os.Args[1]
	var err error
	db, err = sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	q := fmt.Sprintf("SELECT * FROM %s", os.Args[2])
	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.ToUpper(strings.Join(cols, separator)))

	for rows.Next() {
		data := []string{}
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i := range cols {
			data = append(data, columns[i])
		}
		fmt.Println(strings.Join(data, separator))
	}
}
