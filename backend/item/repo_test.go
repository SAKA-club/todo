package item

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SAKA-club/todo/backend/errs"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestGetAll(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock db setup %v", err)
	}
	defer mockDB.Close()

	r := NewRepo(sqlx.NewDb(mockDB, "sqlmock"))

	// Check for an error
	mock.ExpectQuery("SELECT (.+) FROM item WHERE delete_time IS NULL").
		WillReturnError(errs.InternalErr)
	items, err := r.GetAll()
	if len(items) != 0 && err != errs.InternalErr {
		t.Errorf("Expected: %v, Got: %v", "errs.InternalErr", err)
	}

	//	var items []*models.Item
	//	if err := r.db.Select(&items, `SELECT id, title, priority, schedule_time, complete_time FROM item WHERE delete_time IS NULL`); err != nil {
	//		log.Err(err).Msg("could not get items from db")
	//		return []*models.Item{}, errs.InternalErr
	//	}
	//	return items, nil

	// Check for 0 rows
	// TODO:

	// There could be one row
	// There could be more than one row

	//
	//
	//
	//	WillReturnRows(articleMockRows)
	//
	//itemMockRows := sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
	//	AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test")
	//
	//// There could be no rows
	//mock.ExpectExec("INSERT INTO users").
	//	WithArgs("john", AnyTime{}).
	//	WillReturnResult(sqlmock.NewResult(1, 1))
	//
	//
	//_, err = db.Exec("INSERT INTO users(name, created_at) VALUES (?, ?)", "john", time.Now())
	//if err != nil {
	//	t.Errorf("error '%s' was not expected, while inserting a row", err)
	//}
	//
	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
	//
	//mock.ExpectExec("INSERT INTO baskets").WillReturnResult(sqlmock.NewResult(newID, 1))

}
