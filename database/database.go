package database 

import (
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

func CheckTable(db *sql.DB,table string) (bool,error) {
	err := db.QueryRow("SHOW TABLES LIKE ?",table)
	if err != nil {
		return false,err
	}
	return true, nil
}

func CreateTable(db *sql.DB,table string) error {
	_,err := db.Query("CREATE TABLE " + table + "(test INT)")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTable(db *sql.DB,table string) error {
	_,err := db.Query("DROP TABLE " + table)
	if err != nil {
		return err
	}
	return nil
}
