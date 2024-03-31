package jira

import "net/http"

type ResponseHandler[T any] interface {
	Handle(r *http.Response) *T

	OnError(reason string, e error)
}
