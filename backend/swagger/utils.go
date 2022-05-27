package swagger

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"

	"github.com/SAKA-club/todo/backend/errs"
	"github.com/go-openapi/runtime/middleware"
)

func errorHandler(ctx context.Context, label string, err error) middleware.Responder {
	if errs.IsUnauthorized(err) {
		return NewUnauthorized()
	}
	if errs.IsNotFound(err) {
		return NewNotFound()
	}

	var lErr errs.TodoError

	if ok := errors.As(err, &lErr); ok && lErr.Public() {
		l := log.With().Err(err).Str("errorMessage", lErr.Body()).Logger()
		l.WithContext(ctx)

		return NewBadRequest(&lErr)
	}

	log.Ctx(ctx).Error().Err(err).Msg(label)
	return NewInternalServer()
}

func stringErrorHandler(ctx context.Context, label string, err error, payload string) middleware.Responder {
	switch {
	case errs.IsUnauthorized(err):
		return NewUnauthorized()
	case errs.IsNotFound(err):
		return NewNotFound()
	}

	if c, ok := err.(errs.TodoError); ok && c.Public() {
		l := log.With().Err(err).Str("errorMessage", c.Body()).Logger()
		l.WithContext(ctx)
		return NewBadRequestString(c.Body())
	}

	log.Ctx(ctx).Error().Err(err).Msg(label)
	return NewBadRequestString(payload)
}
