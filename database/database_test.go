package database

import (
	"errors"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestDatabaseFunctions(t *testing.T) {
	t.Run("TestGetVersion", func(t *testing.T) {
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
	})

	t.Run("TestCheckTableSuccess", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		defer db.Close()

		table := "hello"

		rows := sqlmock.NewRows([]string{"Tables_in_shop"}).AddRow(table)
		mock.ExpectQuery(regexp.QuoteMeta("SHOW TABLES LIKE '"+table+"'")).WillReturnRows(rows)

		err = CheckTable(db, table)
		if err != nil {
			t.Errorf("Not found table: %s", table)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("TestCheckTableError", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		defer db.Close()

		table := "hello"
		expectedErr := "syntax error near 'LIKE " + "'" + table + "'"
		mock.ExpectQuery("SHOW TABLES LIKE '" + table + "'").WillReturnError(errors.New(expectedErr))

		err = CheckTable(db, table)
		if err == nil {
			t.Errorf("expected error, but got nil")
		} else if err.Error() != expectedErr {
			t.Errorf("expected error: %s, but got: %s", expectedErr, err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("TestCreateTableSuccess", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		defer db.Close()

		table := "test_table"
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE " + table + "(test INT)")).WillReturnResult(sqlmock.NewResult(1, 1))

		err = CreateTable(db, table)
		if err != nil {
			t.Errorf("Error %s", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("TestCreateTableError", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		defer db.Close()

		table := "test"
		expectedErr := "syntax error near 'CREATE TABLE'"

		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE " + table + "(test INT)")).WillReturnError(errors.New(expectedErr))

		err = CreateTable(db, table)
		if err == nil {
			t.Errorf("expected error, but got nil")
		} else if err.Error() != expectedErr {
			t.Errorf("expected error: %s, but got: %s", expectedErr, err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("TestDeleteTableSuccess", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		defer db.Close()

		table := "test"
		mock.ExpectExec(regexp.QuoteMeta("DROP TABLE " + table)).WillReturnResult(sqlmock.NewResult(1, 1))
		err = DeleteTable(db, table)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
