package database

import (
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetVersion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT VERSION()")).WillReturnRows(rows)

	version, err := GetVersion(db)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if version != "8.0.23" {
		t.Errorf("Expected version to be 8.0.23, got %s", version)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestCheckTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Tables_in_shop"}).AddRow("test")
	mock.ExpectQuery(regexp.QuoteMeta("SHOW TABLES LIKE 'test'")).WillReturnRows(rows)

	table := "test"
	found := CheckTable(db, table)
	if !found {
		t.Errorf("Not found table: %s", table)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestCreateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()

	table := "test_table"
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE " + table + "(test INT)")).WillReturnResult(sqlmock.NewResult(1, 1))

	err = CreateTable(db, table)
	if err != nil {
		t.Errorf(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestDeleteTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()

	table := "test"
	mock.ExpectExec(regexp.QuoteMeta("DROP TABLE " + table)).WillReturnResult(sqlmock.NewResult(1, 1))

	err = DeleteTable(db, table)
	if err != nil {
		t.Errorf(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

