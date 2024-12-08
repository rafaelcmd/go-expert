package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rafaelcmd/go-expert/07-apis/configs"
	"github.com/rafaelcmd/go-expert/07-apis/internal/entity"
	"github.com/rafaelcmd/go-expert/07-apis/internal/infra/database"
	"github.com/rafaelcmd/go-expert/07-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Using the Custom LogRequest middleware
	r.Use(LogRequest)
	//The Recoverer middleware is used to recover from panics
	r.Use(middleware.Recoverer)

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.FindAllProducts)
	})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/jwt", userHandler.GetJWT)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}

// LogRequest Middleware function to log requests
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
