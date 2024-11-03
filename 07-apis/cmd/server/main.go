package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rafaelcmd/go-expert/07-apis/configs"
	"github.com/rafaelcmd/go-expert/07-apis/internal/entity"
	"github.com/rafaelcmd/go-expert/07-apis/internal/infra/database"
	"github.com/rafaelcmd/go-expert/07-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	config := configs.LoadConfig(".")
	println(config.DBDriver)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
	}
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.FindAllProducts)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
