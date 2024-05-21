package main

import (
	"net/http"

	"github.com/eron97/project-fullCycle.git/configs"
	"github.com/eron97/project-fullCycle.git/infra/database"
	"github.com/eron97/project-fullCycle.git/infra/webserver/handlers"
	"github.com/eron97/project-fullCycle.git/internal/entity"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.NewConfig()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// productDB é um Ponteiro para Estrutura Product (product_db.go)
	// essa estrutura possui um campo de ponteiro para conexão com db
	// com ela podemos trabalhar de maneira organizada e encapsulada.
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.GettokenAuth(), cfg.GetjwtExpiresIn())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)
	http.ListenAndServe(":8080", r)

}
