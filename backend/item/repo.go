package item

import (
	"database/sql"
	"github.com/SAKA-club/todo/backend/errs"
	"github.com/SAKA-club/todo/backend/gen/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repo {
	return &repo{db: db}

}

func (r repo) GetAll() ([]*models.Item, error) {
	var items []*models.Item
	if err := r.db.Select(&items, `SELECT id, title, priority, schedule_time, complete_time FROM item WHERE delete_time IS NULL`); err != nil {
		log.Err(err).Msg("could not get items from db")
		return []*models.Item{}, errs.InternalErr
	}
	return items, nil
}

func (r repo) Get(ID int64) (*models.Item, error) {
	var item models.Item
	if err := r.db.Get(&item, "SELECT  id, title, priority, schedule_time, complete_time FROM item WHERE id= $1 AND delete_time IS NULL ", ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundErr
		}
		log.Err(err).Msg("could not get by id when searching in database")
		return nil, errs.ErrInternal
	}
	return &item, nil
}

func (r repo) Create(title string, body string, priority bool, scheduleDate time.Time, completeDate time.Time) (*models.Item, error) {
	var item models.Item
	if err := r.db.Get(&item, "INSERT INTO item (title, body, priority, schedule_time, complete_time) VALUES ($1, $2, $3, $4, $5) WHERE delete_time IS NULL RETURNING id, title, priority, schedule_time, complete_time",
		title, body, priority, scheduleDate, completeDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.InputError
		}
		log.Err(err).Msg("could not create item to database")
		return nil, errs.InternalErr
	}

	return &item, nil

}

func (r repo) Delete(ID int64) error {
	_, err := r.db.Exec("DELETE FROM item WHERE id= $1 and delete_time IS NULL", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.ErrNotFound
		}
		log.Err(err).Msg("was not able to delete item or does not exist")
		return errs.ErrInternal
	}
	return nil
}

func (r repo) Update(ID int64, title string, body string, priority bool, scheduleTime time.Time, completeTime time.Time) (*models.Item, error) {
	var item models.Item
	query := "UPDATE item SET title = $1, body= $2, priority = $3, schedule_time = $4, complete_time= $5 WHERE id = $6 AND delete_time IS NULL RETURNING id, title, body, priority, complete_time"
	if err := r.db.Get(&item, query, title, body, priority, scheduleTime, completeTime, ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundErr
		}

		log.Err(err).Msg("could not update database")
		return nil, errs.DBErr
	}

	return &item, nil
}
