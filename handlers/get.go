package handlers

import (
	"net/http"

	"github.com/afterglowflexin/microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponse

// GetProducts returns the list of products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the database
	lp := data.GetProducts()

	//  serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
