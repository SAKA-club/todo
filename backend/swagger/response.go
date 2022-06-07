package swagger

import (
	"net/http"

	"github.com/SAKA-club/todo/backend/errs"
	"github.com/go-openapi/runtime"
)

type BadRequest struct {
	Payload *errs.TodoError `json:"body,omitempty"`
}

func NewBadRequest(p *errs.TodoError) *BadRequest {
	return &BadRequest{Payload: p}
}
func (o *BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusBadRequest)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

type BadRequestString struct {
	Payload string
}

func NewBadRequestString(p string) *BadRequestString {
	return &BadRequestString{Payload: p}
}
func (o *BadRequestString) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusBadRequest)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

type Unauthorized struct {
}

func NewUnauthorized() *Unauthorized {
	return &Unauthorized{}
}

func (o *Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusUnauthorized)
}

type NotFound struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewNotFound() *NotFound {
	return &NotFound{
		Title: "Not found",
		Code:  404,
	}
}

func (o *NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusNotFound)
}

type InternalServer struct {
}

func NewInternalServer() *InternalServer {
	return &InternalServer{}
}

func (o *InternalServer) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusInternalServerError)
}
