package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"working/data"

	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a Products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products from the data base
// responses:
// 		200: productsResponse

// GetProducts handles GET requests and returns all current products.
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products from the datastore.
	lp := data.GetProducts()

	// Convert (serialize) the list to JSON.
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Returns the product with a given id
// responses:
// 		200: productResponse
// 		404: errorResponse

// GetProduct handles GET requests and returns a product.
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	// Returns route variables as a map.
	vars := mux.Vars(r)

	// Get id from route variables.
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle GET Products", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSONInterface(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSONInterface(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSONInterface(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products from the data base
// responses:
// 		201: noContent

// DeleteProduct deletes a product from the data store.
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// Returns route variables as a map.
	vars := mux.Vars(r)

	// Get id from route variables.
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle DELETE Products", id)

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//		200: productResponse
//  	422: errorValidation
//  	501: errorResponse

// AddProduct adds a product to the data store.
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(prod)
}

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//		201: noContent
//  	404: errorResponse
//  	422: errorValidation

// UpdateProducts replaces the Product with a matching id.
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	// Returns route variables as a map.
	vars := mux.Vars(r)

	// Get id from route variables.
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// Update with the new product.
	err = data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
