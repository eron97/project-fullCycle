package main

import (
	"encoding/json"
	"net/http"

	"github.com/eron97/project-fullCycle.git/configs"
	"github.com/eron97/project-fullCycle.git/infra/database"
	"github.com/eron97/project-fullCycle.git/internal/dto"
	"github.com/eron97/project-fullCycle.git/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.NewConfig()
	_ = cfg.GetDBDriver()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// productDB é um Ponteiro para Estrutura Product (product_db.go)
	// essa estrutura possui um campo de ponteiro para conexão com db
	// com ela podemos trabalhar de maneira organizada e encapsulada.
	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)

}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(err)
}
