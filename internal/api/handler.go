package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/sementrof/offerday/internal/api/v1"
	"github.com/sementrof/offerday/internal/middleware"
	_ "github.com/sementrof/prod1/docs"
	"go.uber.org/zap"
)

type TaskServer struct {
	api v1.ApiInterface
}

func NewTaskServer(api v1.ApiInterface) *TaskServer {
	return &TaskServer{api: api}
}

func (s *TaskServer) CreateUsersPost(w http.ResponseWriter, r *http.Request) {
	s.api.CreateUsersPost(w, r)
}

func (s *TaskServer) CreateCategoriesPost(w http.ResponseWriter, r *http.Request) {
	s.api.CreateCategoriesPost(w, r)
}

func (s *TaskServer) CreateEventsPost(w http.ResponseWriter, r *http.Request) {
	s.api.CreateEventsPost(w, r)
}

func (s *TaskServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func SetupRouter(api v1.ApiInterface, logger *zap.Logger) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Use(middleware.LoggingMiddleware(logger))

	server := NewTaskServer(api)

	router.HandleFunc("/", server.HealthCheck).Methods("GET")
	router.HandleFunc("/api/auth/register", server.CreateUsersPost).Methods("POST")
	router.HandleFunc("/api/categories", server.api.CreateCategoriesPost).Methods("POST")
	router.HandleFunc("/api/locations", server.api.CreateLocationsPost).Methods("POST")
	router.HandleFunc("/api/events", server.api.CreateEventsPost).Methods("POST")

	return router

}
