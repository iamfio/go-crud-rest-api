package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamfio/crud-rest-api/repos"
)

type GenericRouter[TT any, T repos.GenericRepository[TT]] struct {
	muxBase    string
	repository *repos.GenericRepository[TT]
}

func (router *GenericRouter[TT, T]) handle(w http.ResponseWriter, r *http.Request) {
	idLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if r.URL.EscapedPath() != router.muxBase && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if idLong != 0 {
			item, err := (*(router.repository)).GetOne(uint(idLong))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("Entity not found!")
				return
			}
			json.NewEncoder(w).Encode(&item)
		} else {
			items := (*(router.repository)).GetList()
			json.NewEncoder(w).Encode(&items)
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		var model TT
		json.NewDecoder(r.Body).Decode(&model)
		(*(router.repository)).Create(model)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&model)
	case http.MethodPut:
		var model TT
		json.NewDecoder(r.Body).Decode(&model)
		_, err := (*(router.repository)).Update(uint(idLong), model)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Entity not found!")
			return
		}
		json.NewEncoder(w).Encode(&model)
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if idLong != 0 {
			_, err := (*(router.repository)).DeleteOne(uint(idLong))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("Entity not found!")
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (router *GenericRouter[TT, T]) registerRoutes(mux *mux.Router) {
	mux.HandleFunc(router.muxBase, router.handle)
	mux.HandleFunc(fmt.Sprintf("%v/{id}", router.muxBase), router.handle)
}

func NewGenericRouter[TT any, T repos.GenericRepository[TT]](muxBase string, mux *mux.Router, repository *repos.GenericRepository[TT]) *GenericRouter[TT, T] {
	router := GenericRouter[TT, T]{
		muxBase:    muxBase,
		repository: repository,
	}
	router.registerRoutes(mux)
	return &router
}
