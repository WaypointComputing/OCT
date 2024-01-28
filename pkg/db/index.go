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
func (db *DB) ExecuteSQL(filename string) error {
	path := path.Join(db_dir, sql_dir, filename+".sql")

	out, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sql := string(out)

	_, err = db.db.Exec(sql)
	return err
}

func SetupDB(url string) error {
	var dbExisted bool = false
	if _, err := os.Stat(url); errors.Is(err, os.ErrNotExist) {
		dbExisted = true
	}

	db, err := sql.Open("libsql", "file:"+url)
	if err != nil {
		return err
	}

	Db = &DB{
		db,
	}

	if dbExisted {
		seedData()
	}

	return nil
}

func seedData() error {
	err := Db.ExecuteSQL("WaypointDBCreation")
	if err != nil {
		return err
	}

	utils.Log("DB: Initialized Schema!")

	err = Db.ExecuteSQL("SampleData")
	if err != nil {
		return err
	}

	utils.Log("DB: Inserted Sample Data!")

	return nil
}
