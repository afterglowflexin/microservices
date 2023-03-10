package handlers

import (
	"net/http"

	"github.com/afterglowflexin/microservices/data"
)

// test data for POST - "{\"id\": 1,\"name\":\"Tea\",\"description\":\"Hot cup of tea\"}"

// swagger:route POST /products products addProduct
// Adds a product to a product list
// responses:
// 201: noContent
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
