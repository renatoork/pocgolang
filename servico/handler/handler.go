package handler

import (
	"encoding/json"
	"net/http"
)

func Decode(v interface{}, req *http.Request, res http.ResponseWriter) bool {
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

type Servico interface {
	Insere() interface{}
}
