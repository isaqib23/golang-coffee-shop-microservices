package user

import (
	"encoding/json"
	"github.com/isaqib23/golang-coffee-shop-microservices/domain"
	"net/http"
)

type handler struct {
	svc domain.UserService
}

func (h *handler) FetchByName() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		name := "example" // replace "example" with the actual name you want to fetch

		result, err := h.svc.FetchByName(nil, name)
		if err != nil {
			// Handle the error here, e.g. write an error response to the response writer
			// and return from the function
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the result to the response writer
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(result)
	}
}
