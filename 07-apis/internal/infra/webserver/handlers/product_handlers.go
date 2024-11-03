package handlers

import (
	"encoding/json"
	"github.com/rafaelcmd/go-expert/07-apis/internal/dto"
	"github.com/rafaelcmd/go-expert/07-apis/internal/entity"
	"github.com/rafaelcmd/go-expert/07-apis/internal/infra/database"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductDTO dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&createProductDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(createProductDTO.Name, createProductDTO.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
