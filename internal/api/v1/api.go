package ф

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/sementrof/offerday/internal/db"
	"github.com/sementrof/offerday/internal/deps"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type ApiInterface interface {
	CreateUsersPost(w http.ResponseWriter, r *http.Request)
	CreateCategoriesPost(w http.ResponseWriter, r *http.Request)
	CreateLocationsPost(w http.ResponseWriter, r *http.Request)
	CreateEventsPost(w http.ResponseWriter, r *http.Request)
}

type ApiImplemented struct {
	deps *deps.Dependencies
}

func NewApi(deps *deps.Dependencies) *ApiImplemented {
	return &ApiImplemented{
		deps: deps,
	}
}

func (im *ApiImplemented) CreateUsersPost(w http.ResponseWriter, r *http.Request) {
	var inputUsers db.Users
	ctx := context.Background()
	if err := json.NewDecoder(r.Body).Decode(&inputUsers); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(inputUsers); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputUsers.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user := db.Users{
		Id:        inputUsers.Id,
		Name:      inputUsers.Name,
		Password:  string(hashedPassword),
		Email:     inputUsers.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	errs := im.deps.DB.Users.NewInsertUser(ctx, &user)
	if errs != nil {
		im.deps.Logger.Error("Failed to create user", zap.Error(errs))
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (im *ApiImplemented) CreateCategoriesPost(w http.ResponseWriter, r *http.Request) {
	var inputCategories db.Categories
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&inputCategories); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(inputCategories); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	categories := db.Categories{
		Id:        inputCategories.Id,
		Name:      inputCategories.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	errs := im.deps.DB.Categories.NewInsertCategories(ctx, &categories)
	if errs != nil {
		im.deps.Logger.Error("Failed to create user", zap.Error(errs))
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (im *ApiImplemented) CreateLocationsPost(w http.ResponseWriter, r *http.Request) {
	var inputLocations db.Locations
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&inputLocations); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(inputLocations); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	locations := db.Locations{
		Id:        inputLocations.Id,
		Name:      inputLocations.Name,
		Addres:    inputLocations.Addres,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	errs := im.deps.DB.Locations.NewInsertLocations(ctx, &locations)
	if errs != nil {
		im.deps.Logger.Error("Failed to create user", zap.Error(errs))
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (im *ApiImplemented) CreateEventsPost(w http.ResponseWriter, r *http.Request) {
	var inputevents db.Events
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&inputevents); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(inputevents); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	events := db.Events{
		Id:          inputevents.Id,
		Title:       inputevents.Title,
		Description: inputevents.Description,
		Date:        inputevents.Date,
		OrganizerId: inputevents.OrganizerId,
		CategoryId:  inputevents.CategoryId,
		LocationId:  inputevents.LocationId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	errs := im.deps.DB.Events.NewInsertEvents(ctx, &events)
	if errs != nil {
		im.deps.Logger.Error("Failed to create user", zap.Error(errs))
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
