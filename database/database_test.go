package database

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)


//func TestDatabaseConnect(t *testing.T) {
//	err := DatabaseConnect(testConnectionString)
//	if err != nil {
//		t.Fatalf("Expected no error, got %v", err)
//	}
//}

func TestGetVersion(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23")
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(rows)

	version,err = GetVersion(db)
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
	:defer db.Close()

	
}

func TestCreateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()
}

func TestDeleteTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	defer db.Close()
}

