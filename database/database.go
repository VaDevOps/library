package database 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
	err error
	version string
)

func DatabaseConnect(connection string) error {
	DB, err = sql.Open("mysql", connection)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func GetVersion() (string,error) {
	err = DB.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "",err
	}
	return version,nil
}

func CheckTable(table string) (bool,error) {
	res, err := DB.Query("SHOW TABLES LIKE ?",table)
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

func CreateTable(table string) error {
	_,err := DB.Query("CREATE TABLE " + table + "(test INT)")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTable(table string) error {
	_,err := DB.Query("DROP TABLE " + table)
	if err != nil {
		return err
	}
	return nil
}
