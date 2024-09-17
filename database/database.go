package database 

import (
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect(connection string) (*sql.DB,error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		return nil,err
	}
	return db,nil
}

func GetVersion(db *sql.DB) (string,error) {
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "",err
	}
	return version,nil
}

func CheckTable(db *sql.DB,table string) error {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'",table)
	rows,err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return errors.New("table not found")
}

func CreateTable(db *sql.DB,table string) error {
	_,err := db.Exec("CREATE TABLE " + table + "(test INT)")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTable(db *sql.DB,table string) error {
	_,err := db.Exec("DROP TABLE " + table)
	if err != nil {
		return err
	}
	return nil
}
