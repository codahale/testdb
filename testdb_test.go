package testdb_test

import (
	"database/sql"
	"fmt"

	"github.com/codahale/testdb"
	_ "github.com/go-sql-driver/mysql"
)

func Example() {
	baseDSN := testdb.Env("MYSQL_DB", testdb.LocalMySQL)

	tdb, err := testdb.Open("mysql", baseDSN)
	if err != nil {
		panic(err)
	}
	defer tdb.Close()

	db, err := sql.Open("mysql", baseDSN+tdb.Name())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE things(id INT AUTO_INCREMENT, PRIMARY KEY (id))`); err != nil {
		panic(err)
	}

	var count int
	if err := db.QueryRow(`SELECT COUNT(id) FROM things`).Scan(&count); err != nil {
		panic(err)
	}

	fmt.Println(count)

	if _, err := db.Exec(`INSERT INTO things () VALUES ()`); err != nil {
		panic(err)
	}

	var newCount int
	if err := db.QueryRow(`SELECT COUNT(id) FROM things`).Scan(&newCount); err != nil {
		panic(err)
	}

	fmt.Println(newCount)

	// Output:
	// 0
	// 1
}
