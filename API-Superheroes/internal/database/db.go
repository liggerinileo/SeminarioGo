package database

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	"github.com/liggerinileo/SeminarioGo/API-Superheroes/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

// PopulateDatabase ...
func PopulateDatabase(db *sqlx.DB, filepath string) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}

	query := string(file)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// CreateSchema ...
func CreateSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS superheroe (
		id integer primary key autoincrement,
		name varchar,
		speciality varchar,
		damage varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
