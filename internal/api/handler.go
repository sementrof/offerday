package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/sementrof/offerday/internal/api/v1"
	_ "github.com/sementrof/prod1/docs"
)

type TaskServer struct {
	api v1.ApiInterface
}

func NewTaskServer(api v1.ApiInterface) *TaskServer {
	return &TaskServer{api: api}
}

func (s *TaskServer) CreateOrganizationPost(w http.ResponseWriter, r *http.Request) {
	s.api.CreateOrganizationPost(w, r)
}

func SetupRouter(api v1.ApiInterface) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	// server := NewTaskServer(api)
	// router.HandleFunc("/registration/user", server.createUserPost).Methods("POST")

	return router

}
