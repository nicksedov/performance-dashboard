package handler

import (
	"net/http"
)

type ResponseHandler[T any] interface {

	Handle(r *http.Response, dto *T) 

	OnError(reason string, e error)
}
