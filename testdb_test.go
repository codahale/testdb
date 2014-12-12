package testdb_test

import (
	"database/sql"
	"fmt"

	"github.com/codahale/testdb"
	_ "github.com/lib/pq"
)

func Example() {
	baseDSN := testdb.Env("PG_DB", "sslmode=disable")

	tdb, err := testdb.Open("postgres", baseDSN+" dbname=postgres")
	if err != nil {
		panic(err)
	}
	defer tdb.Close()

	db, err := sql.Open("postgres", baseDSN+" dbname="+tdb.Name())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE things(id SERIAL, PRIMARY KEY (id))`); err != nil {
		panic(err)
	}

	var count int
	if err := db.QueryRow(`SELECT COUNT(id) FROM things`).Scan(&count); err != nil {
		panic(err)
	}

	fmt.Println(count)

	if _, err := db.Exec(`INSERT INTO things DEFAULT VALUES`); err != nil {
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
