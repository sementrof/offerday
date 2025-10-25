package ф

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/sementrof/offerday/internal/deps"
)

var validate = validator.New()

type ApiInterface interface {
	CreateOrganizationPost(w http.ResponseWriter, r *http.Request)
}

type ApiImplemented struct {
	deps *deps.Dependencies
}

func NewApi(deps *deps.Dependencies) *ApiImplemented {
	return &ApiImplemented{
		deps: deps,
	}
}

func (im *ApiImplemented) CreateOrganizationPost(w http.ResponseWriter, r *http.Request) {
	// var input models.OrganizationInput

}
