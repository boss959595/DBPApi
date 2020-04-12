package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCon() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	checkErr(err)

	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
