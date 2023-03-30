package handlers

import (
	"github.com/isaqib23/golang-coffee-shop-microservices/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, *r)
		return
	}

	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r http.Request) {
	// fetch the products from datastore
	products := data.GetProducts()

	// serialize from list to JSON
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Error while getting products", http.StatusInternalServerError)
	}
}
