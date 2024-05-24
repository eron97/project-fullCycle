package main

import (
	"net/http"

	"github.com/eron97/project-fullCycle.git/configs"
	_ "github.com/eron97/project-fullCycle.git/docs"
	"github.com/eron97/project-fullCycle.git/internal/entity"
	"github.com/eron97/project-fullCycle.git/internal/infra/database"
	"github.com/eron97/project-fullCycle.git/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Experct API Example
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Eron Betine

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApikeyAuth
// @in header
// @name Authorization

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
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(300))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.GettokenAuth()))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)

}
