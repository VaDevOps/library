package database 

import (
	"fmt"
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
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "",err
	}
	return version,nil
}

func CheckTable(db *sql.DB,table string) (bool,error) {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)
	res, err := DB.Query(query)
	if err != nil {
		return false,err
	}
	defer res.Close()

	var foundtable string
	if res.Next() {
		if err := res.Scan(&foundtable); err != nil {
			return false,err
		}
		return true, nil //found
	}
	return false, nil // not found
}

func CreateTable(db *sql.DB,table string) error {
	_,err := DB.Query("CREATE TABLE " + table + "(test INT)")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTable(db *sql.DB,table string) error {
	_,err := DB.Query("DROP TABLE " + table)
	if err != nil {
		return err
	}
	return nil
}
