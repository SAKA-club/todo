package swagger

import (
	"context"
	"github.com/SAKA-club/todo/backend/errs"
	"github.com/SAKA-club/todo/backend/gen/models"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations"
	"github.com/SAKA-club/todo/backend/gen/restapi/operations/item"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
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
			log.Err(err).Msg("Get all items handler")

			return item.NewGetAllInternalServerError().WithPayload(&errs.ErrInternal)
		}
		return item.NewGetAllOK().WithPayload(result)

	})

	api.ItemGetHandler = item.GetHandlerFunc(func(params item.GetParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		result, err := service.Get(ctx, params.ID)
		if err != nil {
			if errs.IsNotFound(err) {
				return item.NewGetNotFound().WithPayload(&errs.ErrNotFound)
			}
			log.Err(err).Msg("Get item handler")
			return item.NewGetInternalServerError().WithPayload(&errs.ErrInternal)
		}
		return item.NewGetOK().WithPayload(result)
	})

	api.ItemCreateHandler = item.CreateHandlerFunc(func(params item.CreateParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		i := params.Body
		if i == nil || i.Title == nil || *i.Title == "" {
			return item.NewCreateBadRequest().WithPayload(&errs.ErrInvalidRequest)
		}
		result, err := service.Create(ctx, *i.Title, i.Body, i.Priority, time.Time(i.ScheduleTime), time.Time(i.CompleteTime))
		if err != nil {
			log.Err(err).Msg("put create items handler")
			return item.NewCreateInternalServerError().WithPayload(&errs.ErrInternal)

		}

		return item.NewCreateCreated().WithPayload(result)

	})

	api.ItemDeleteHandler = item.DeleteHandlerFunc(func(params item.DeleteParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		err := service.Delete(ctx, params.ID)
		if err != nil {
			if errs.IsNotFound(err) {
				return item.NewDeleteNotFound().WithPayload(&errs.ErrNotFound)
			}
			log.Error().Err(err).Msg("Item Delete Handler")
			return item.NewDeleteInternalServerError().WithPayload(&errs.ErrInternal)
		}

		return item.NewDeleteNoContent()

	})

	api.ItemUpdateHandler = item.UpdateHandlerFunc(func(params item.UpdateParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		i := params.Body
		if i == nil || i.Title == nil || *i.Title == "" {
			return item.NewUpdateBadRequest().WithPayload(&errs.ErrInvalidRequest)
		}

		result, err := service.Update(ctx, i.ID, *i.Title, i.Body, i.Priority, time.Time(i.ScheduleTime), time.Time(i.CompleteTime))

		if err != nil {
			if errs.IsNotFound(err) {
				return item.NewUpdateNotFound().WithPayload(&errs.ErrNotFound)
			}
			log.Err(err).Msg("put update item handler")
			return item.NewUpdateInternalServerError().WithPayload(&errs.ErrInternal)
		}

		return item.NewUpdateOK().WithPayload(result)
	})
}
