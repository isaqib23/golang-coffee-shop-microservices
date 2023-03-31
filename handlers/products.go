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

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// fetch the products from datastore
	products := data.GetProducts()

	// serialize from list to JSON
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Error while getting products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Invalid Data", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}
