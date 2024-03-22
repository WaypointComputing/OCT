package db

import (
	"database/sql"
	"errors"
	"os"
	"path"
	"waypoint/pkg/utils"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const db_dir string = "db"
const sql_dir string = "sql"

var Db *sql.DB

// Executes an .sql file from the project_root/db/sql directory.
// Takes the name of the file (without the file extension) as input.
func ExecuteSQL(filename string, args ...interface{}) (sql.Result, error) {
	sql, err := readSQLFile(filename)
	if err != nil {
		return nil, err
	}
	return Db.Exec(sql, args...)
}

func QuerySQL(filename string, args ...interface{}) (*sql.Rows, error) {
	sql, err := readSQLFile(filename)
	if err != nil {
		return nil, err
	}
	return Db.Query(sql, args...)
}

func QueryRowSQL(filename string, args ...interface{}) (*sql.Row, error) {
	sql, err := readSQLFile(filename)
	if err != nil {
		return nil, err
	}
	return Db.QueryRow(sql, args...), nil
}

func readSQLFile(filename string) (string, error) {
	path := path.Join(db_dir, sql_dir, filename+".sql")
	out, err := os.ReadFile(path)
	return string(out), err
}

func SetupDB(url string) error {
	var dbExisted bool = true
	if _, err := os.Stat(url); errors.Is(err, os.ErrNotExist) {
		dbExisted = false
	}

	db, err := sql.Open("libsql", "file:"+url)
	if err != nil {
		return err
	}

	Db = db

	if !dbExisted {
		seedData()
	}

	return nil
}

func seedData() error {
	_, err := ExecuteSQL("WaypointDBCreation")
	if err != nil {
		return err
	}

	utils.Log("DB: Initialized Schema!")

	_, err = ExecuteSQL("SampleData")
	if err != nil {
		return err
	}

	utils.Log("DB: Inserted Sample Data!")

	return nil
}
