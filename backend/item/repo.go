package item

import (
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
	if err := r.db.Select(&items, `SELECT id, title, priority, schedule_time, complete_time FROM item`); err != nil {
		log.Err(err).Msg("could not get items from db")
		return []*models.Item{}, err
	}
	return items, nil
}

func (r repo) Get(ID int64) (*models.Item, error) {
	var item models.Item
	if err := r.db.Get(&item, "SELECT  id, title, priority, schedule_time, complete_time FROM item WHERE id= $1", ID); err != nil {
		log.Err(err).Msg("could not find by id")
		return nil, err
	}
	return &item, nil
}

func (r repo) Create(title string, body string, priority bool, scheduleDate time.Time, completeDate time.Time) (*models.Item, error) {
	var item models.Item
	if err := r.db.Get(&item, "INSERT INTO item (title, body, priority, schedule_time, complete_time) VALUES ($1, $2,$3, $4,$5) RETURNING id, title, priority, schedule_time, complete_time",
		title, body, priority, scheduleDate, completeDate); err != nil {
		log.Err(err).Msg("could not add to database")
		return nil, err
	}

	return &item, nil

}

func (r repo) Delete(ID int64) error {
	_, err := r.db.Exec("DELETE FROM item WHERE id= $1", ID)
	if err != nil {
		log.Err(err).Msg("was not able to delete item or does not exist")
	}
	return nil
}

func (r repo) Update(ID int64, title string, body string, priority bool, scheduleTime time.Time, completeTime time.Time) (*models.Item, error) {
	var item models.Item
	if err := r.db.Get(&item, "UPDATE item SET title =$1, body= $2, priority = $3, schedule_time =$4, complete_time=$5 WHERE id = $6  RETURNING id, title, body, priority, complete_time",
		title, body, priority, scheduleTime, completeTime, ID); err != nil {
		log.Err(err).Msg("could not add to database")
		return nil, err
	}

	return &item, nil
}
