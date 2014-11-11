// Package testdb provides an easy way to, given a running database server,
// create temporary, isolated databases to use in testing.
package testdb

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"os"
)

// TestDB is a temporary database for testing.
type TestDB struct {
	name string

	db *sql.DB
}

// Open creates a new TestDB given the driver name and the data source name.
func Open(driverName, dataSourceName string) (*TestDB, error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}

	name := fmt.Sprintf("tmpdb_%016x", buf)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("CREATE DATABASE " + name); err != nil {
		return nil, err
	}

	return &TestDB{name: name, db: db}, nil
}

// Name gets the name of the created test database
func (tdb *TestDB) Name() string {
	return tdb.name
}

// Close drops the testing databases and disconnects from the server.
//
// N.B.: Not calling this will result in your database server accumulating many
// databases.
func (tdb *TestDB) Close() error {
	if _, err := tdb.db.Exec("DROP DATABASE " + tdb.name); err != nil {
		return err
	}
	return tdb.db.Close()
}

// Env returns an environment variable, or if the environment variable is not
// present, the default value.
func Env(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

const (
	// LocalMySQL is the DSN for a basic install of MySQL accepting
	// unauthenticated connections for the root user on the default Unix socket.
	LocalMySQL = "root:@unix()/"
)
