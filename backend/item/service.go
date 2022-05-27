package item

import (
	"context"
	"github.com/SAKA-club/todo/backend/gen/models"
	"github.com/go-openapi/strfmt"
)

type service struct {
	r *repo
}

func NewService(r *repo) *service {
	return &service{r: r}

}

func (s service) GetAll(ctx context.Context) ([]*models.Item, error) {
	return s.r.GetAll()

}

func (s service) Get(ctx context.Context, ID int64) (*models.Item, error) {
	return s.r.Get(ID)

}

func (s service) Create(ctx context.Context, title string, body string, priority bool, scheduleDate strfmt.Date, completeDate strfmt.Date) (*models.Item, error) {
	return s.r.Create(title, body, priority, scheduleDate, completeDate)
}

func (s service) Delete(ctx context.Context, ID int64) error {
	return s.r.Delete(ID)

}

func (s service) Update(ctx context.Context, ID int64, title string, body string, priority bool, scheduleTime strfmt.Date, completeTime strfmt.Date) (*models.Item, error) {
	return s.r.Update(ID, title, body, priority, scheduleTime, completeTime)

}
