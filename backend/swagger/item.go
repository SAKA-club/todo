package swagger

import (
	"context"
	"github.com/SAKA-club/todo/backend/errs"
	"github.com/SAKA-club/todo/backend/gen/models"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations/item"
	"github.com/go-openapi/runtime/middleware"
	"time"
)

type ItemService interface {
	GetAll(ctx context.Context) ([]*models.Item, error)
	Get(ctx context.Context, ID int64) (*models.Item, error)
	Create(ctx context.Context, title string, body string, priority bool, scheduleTime time.Time, completeTime time.Time) (*models.Item, error)
	Delete(ctx context.Context, ID int64) error
	Update(ctx context.Context, ID int64, title string, body string, priority bool, scheduleTime time.Time, completeTime time.Time) (*models.Item, error)
}

func Item(api *operations.TodoAPI, service ItemService) {
	api.ItemGetAllHandler = item.GetAllHandlerFunc(func(params item.GetAllParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		result, err := service.GetAll(ctx)
		if err != nil {
			return item.NewGetNotFound().WithPayload(err)
		}
		return item.NewGetAllOK().WithPayload(result)

	})

	api.ItemGetHandler = item.GetHandlerFunc(func(params item.GetParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		result, err := service.Get(ctx, params.ID)
		if err != nil {
			return item.NewGetNotFound().WithPayload(err.Error())
		}
		return item.NewGetOK().WithPayload(result)
	})

	api.ItemCreateHandler = item.CreateHandlerFunc(func(params item.CreateParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		i := params.Body
		result, err := service.Create(ctx, *i.Title, i.Body, i.Priority, time.Time(i.ScheduleTime), time.Time(i.CompleteTime))
		if err != nil {
			return item.NewCreateBadRequest().WithPayload(err)
		}

		return item.NewCreateCreated().WithPayload(result)

	})

	api.ItemDeleteHandler = item.DeleteHandlerFunc(func(params item.DeleteParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		err := service.Delete(ctx, params.ID)
		if err != nil {
			if errs.IsNotFound(err) {
				return item.NewUpdateNotFound().WithPayload(err)
			}
			return item.NewDeleteBadRequest().WithPayload(err)
		}

		return item.NewUpdateNotFound().WithPayload(err)

	})

	api.ItemUpdateHandler = item.UpdateHandlerFunc(func(params item.UpdateParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		i := params.Body

		result, err := service.Update(ctx, i.ID, *i.Title, i.Body, i.Priority, time.Time(i.ScheduleTime), time.Time(i.CompleteTime))

		if err != nil {
			if errs.IsNotFound(err) {
				return item.NewUpdateNotFound().WithPayload(err)
			}
			if errs.IsNotNoRows(err) {
				return item.NewCreateBadRequest().WithPayload(err)
			}
		}

		return item.NewUpdateOK().WithPayload(result)

	})

}
