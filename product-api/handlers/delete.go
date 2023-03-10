package handlers

import (
	"net/http"
	"strconv"

	"github.com/afterglowflexin/microservices/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
//
// responses:
// 	201: noContent
//  404: errorResponse
//  501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
