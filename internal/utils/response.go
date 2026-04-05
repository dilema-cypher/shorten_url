package utils

import "net/http"

type ErrorSetter interface {
	SetError(err string)
}

func SetErrorOnResponse(w http.ResponseWriter, err string) {
	if es, ok := w.(ErrorSetter); ok {
		es.SetError(err)
	}
}
