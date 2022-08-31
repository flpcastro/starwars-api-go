package utils

import (
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, err error) {
	JsonRes(w, statusCode, struct {
		Erro string
	}{
		Erro: err.Error(),
	})
}
