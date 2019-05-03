package servico

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Params map[string]string

type Servico interface {
	Insert(vars Params) (interface{}, error)
	Get(vars Params) (interface{}, error)
	Update(vars Params) (interface{}, error)
}

func Handler(res http.ResponseWriter, req *http.Request, obj Servico) {
	vars := mux.Vars(req)
	switch req.Method {
	case "POST":
		if err := json.NewDecoder(req.Body).Decode(&obj); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		newObj, err := obj.Insert(vars)
		if err == nil {
			res.Header().Set("content-type", "application/json")
			json.NewEncoder(res).Encode(newObj)
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	case "GET":
		obj, err := obj.Get(vars)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(res).Encode(obj); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		if err := json.NewDecoder(req.Body).Decode(&obj); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		alterado, err := obj.Update(vars)
		if err == nil || err.Error() == "" {
			res.Header().Set("content-type", "application/json")
			json.NewEncoder(res).Encode(alterado)
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

	}

}
