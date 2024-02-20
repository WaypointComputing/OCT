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

type DB struct {
	db *sql.DB
}

var Db *DB

// Executes an .sql file from the project_root/db/sql directory.
// Takes the name of the file (without the file extension) as input.
func (db *DB) ExecuteSQL(filename string) (sql.Result, error) {
	path := path.Join(db_dir, sql_dir, filename+".sql")

	out, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sql := string(out)

	return db.db.Exec(sql)
}

func (db *DB) QuerySQL(filename string) (*sql.Rows, error) {
	path := path.Join(db_dir, sql_dir, filename+".sql")

	out, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sql := string(out)

	return db.db.Query(sql)
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

	Db = &DB{
		db,
	}

	if !dbExisted {
		seedData()
	}

	return nil
}

func seedData() error {
	_, err := Db.ExecuteSQL("WaypointDBCreation")
	if err != nil {
		return err
	}

	utils.Log("DB: Initialized Schema!")

	_, err = Db.ExecuteSQL("SampleData")
	if err != nil {
		return err
	}

	utils.Log("DB: Inserted Sample Data!")

	return nil
}
