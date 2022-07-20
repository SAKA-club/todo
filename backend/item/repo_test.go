package item

import (
	"database/sql"
	_ "database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SAKA-club/todo/backend/errs"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
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
	if len(items) != 0 || err != errs.InternalErr {
		t.Errorf("Expected: %v, Got: %v", "errs.InternalErr", err)
	}

	// Check for 0 rows
	rows := sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"})
	mock.ExpectQuery("SELECT (.+) FROM item WHERE delete_time IS NULL").WillReturnRows(rows)
	items, err = r.GetAll()
	if len(items) != 0 || err != nil {
		t.Errorf("Expected %v, Got %v", 0, len(items))
	}

	// There could be one row
	rows = sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"}).
		AddRow(1, "Test 1", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")
	mock.ExpectQuery("SELECT (.+) FROM item WHERE delete_time IS NULL").WillReturnRows(rows)
	items, err = r.GetAll()

	if len(items) != 1 || err != nil {
		t.Errorf("Expected 1 row, got %v", len(items))
	}

	// There could be more than one row
	rows = sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"}).
		AddRow(1, "Test 1 ", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05").
		AddRow(2, "Test 2", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")

	mock.ExpectQuery("SELECT (.+) FROM item WHERE delete_time IS NULL").WillReturnRows(rows)
	items, err = r.GetAll()
	if len(items) != 2 || err != nil {
		t.Errorf("Expected 2 or more rows, got %v", len(items))
	}
}

func TestGet(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock db setup %v", err)
	}
	defer mockDB.Close()

	r := NewRepo(sqlx.NewDb(mockDB, "sqlmock"))

	//checking for error
	mock.ExpectQuery("SELECT id, title, priority, schedule_time, complete_time FROM item WHERE id= (.+) AND delete_time IS NULL").
		WillReturnError(sql.ErrNoRows)

	item, err := r.Get(1)
	if item != nil || err != errs.NotFoundErr {
		t.Errorf("Expected: %v, Got: %v", "errs.InternalErr", err)
	}
	// check if item is not there and you try to retrieve it?
	rows := sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"})
	mock.ExpectQuery("SELECT id, title, priority, schedule_time, complete_time FROM item WHERE id= (.+) AND delete_time IS NULL").WillReturnRows(rows)
	item, err = r.Get(1)
	if item != nil || err != nil {
		t.Errorf("Expected item to be nil, Got %v", item)
	}

	// check if you can get item based on id

	rows = sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"}).
		AddRow(1, "Test 1", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")

	mock.ExpectQuery("SELECT id, title, priority, schedule_time, complete_time FROM item WHERE id= (.+) AND delete_time IS NULL").WillReturnRows(rows)

	item, err = r.Get(1)
	if item == nil || err != nil {
		t.Errorf("Expected item to return item{1, Test 1, true, 2006-01-02 15:04:05, 2006-01-02 15:04:05} got %v", item)
	}

	// check if the item is deleted and you cannot retrive it

	rows = sqlmock.NewRows([]string{"id", "title", "priority", "schedule_time", "complete_time"}).
		AddRow(1, "Test 1", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")

	mock.ExpectQuery("SELECT id, title, priority, schedule_time, complete_time FROM item WHERE id= (.+) AND delete_time IS NOT NULL").WillReturnRows(rows)

	item, err = r.Get(1)
	if item == nil || err != nil {
		t.Errorf("Expected item to return item{1, Test 1, true, 2006-01-02 15:04:05, 2006-01-02 15:04:05} got %v", item)
	}

}

func TestCreate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock db setup %v", err)
	}
	defer mockDB.Close()

	r := NewRepo(sqlx.NewDb(mockDB, "sqlmock"))

	//checking for error
	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnError(sql.ErrNoRows)

	t1 := time.Now()
	item, err := r.Create("Test 1", "Hello this is a test", true, &t1, &t1)

	if item != nil || err != errs.InputError {
		t.Errorf("Expected: %v, Got: %v", "errs.InputError", err)
	}

	//checking for everything filled out
	rows := sqlmock.NewRows([]string{"id", "title", "body", "priority", "schedule_time", "complete_time"}).
		AddRow(1, "Test 2", "hello this is a test", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")

	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnRows(rows)

	item, err = r.Create("Test 2", "Hello this is a test", true, &t1, &t1)
	if item == nil || err != nil {
		t.Errorf("expected item to return item {id 1, test 2, hello this is a test, true, %v, %v} Got: %v", t1, t1, item)
	}

	//checking for everything null except for title
	rows = sqlmock.NewRows([]string{"id", "title", "body", "priority", "schedule_time", "complete_time"}).
		AddRow(2, "Test 3", " ", false, nil, nil)
	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnRows(rows)
	item, err = r.Create("Test 3", "", false, nil, nil)
	if item == nil || err != nil {
		t.Errorf("expected item to return item {id 2, test 3, false, nil, nil} Got: %v", item)
	}

	//checking for body to be empty/nil
	rows = sqlmock.NewRows([]string{"id", "title", "body", "priority", "schedule_time", "complete_time"}).
		AddRow(3, "Test 4", " ", true, "2006-01-02 15:04:05", "2006-01-02 15:04:05")

	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnRows(rows)
	item, err = r.Create("Test 4", "", true, &t1, &t1)
	if item == nil || err != nil {
		t.Errorf("expected item to return item {id 2, test 3,  ' ', true, %v, %v} Got: %v", t1, t1, item)
	}

	//checking for schedule time to be empty/nil
	rows = sqlmock.NewRows([]string{"id", "title", "body", "priority", "schedule_time", "complete_time"}).
		AddRow(4, "Test 5", " this is test number 5 ", true, nil, "2006-01-02 15:04:05")

	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnRows(rows)
	item, err = r.Create("Test 5", "this is test number 5", true, nil, &t1)
	if item == nil || err != nil {
		t.Errorf("expected item to return item {id 2, test 3,  ' ', true, nil, %v} Got: %v", t1, item)
	}

	//checking for complete time to be empty/nil
	rows = sqlmock.NewRows([]string{"id", "title", "body", "priority", "schedule_time", "complete_time"}).
		AddRow(5, "Test 6", " this is test number 6 ", true, "2006-01-02 15:04:05", nil)

	mock.ExpectQuery("INSERT INTO item (.+)" +
		" VALUES (.+) RETURNING id, title, body, priority, schedule_time, complete_time").
		WillReturnRows(rows)
	item, err = r.Create("Test 5", "this is test number 6", true, &t1, nil)
	if item == nil || err != nil {
		t.Errorf("expected item to return item {id 2, test 3,  ' ', true, %v, nil} Got: %v", t1, item)
	}
}
