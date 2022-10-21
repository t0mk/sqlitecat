package main

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var ()

func OpenDB(dbfile string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	fileName := os.Args[1]
	db, err := OpenDB(fileName)
	if err != nil {
		panic(err)
	}
	tabs, err := db.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	tab := tabs[0]
	if len(tabs) > 1 {
		fmt.Println("More than 1 tab, using arg[2]")
		tab = os.Args[2]
	}

	cts, err := db.Migrator().ColumnTypes(tab)
	if err != nil {
		panic(err)
	}
	typs := []string{}
	ctypesmap := map[string]string{}
	for _, ct := range cts {
		typs = append(typs, ct.Name())
		cty, _ := ct.ColumnType()
		ctypesmap[ct.Name()] = cty
	}
	sep := os.Getenv("SEP")
	if sep == "" {
		sep = " | "
	}
	uq := os.Getenv("QUERY")
	q := fmt.Sprintf("select * from %s", tab)
	if uq != "" {
		q = fmt.Sprintf("select * from %s where %s", tab, uq)
	}

	var result []map[string]interface{}

	tx := db.Raw(q).Scan(&result)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	colstring := strings.Join(typs, sep)
	fmt.Println(colstring)
	for _, v := range result {
		for j, c := range typs {
			fmt.Printf("%v", v[c])
			if j < (len(typs) - 1) {
				fmt.Printf("%s", sep)
			}
		}
		fmt.Println()
	}
	fmt.Println(colstring)
}
