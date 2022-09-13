package handlers

import (
	"net/http"

	"github.com/vitorcarra/product-api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
