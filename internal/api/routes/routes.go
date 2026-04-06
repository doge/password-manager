package routes

import (
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type routes struct {
	router *chi.Mux
}

func InitializeRoutes() (*chi.Mux, error) {
	api := routes{
		router: chi.NewMux(),
	}

	// Middleware

	return api.router, nil
}

func AuthRoutes(router *chi.Mux, store *mongo.Database) error {

	// Initialize repository
	// Initalize service
	// Initialize controller

	router.Route("/auth", func(authRouter chi.Router) {
		// authRouter.Post("/login", )
		// authRouter.Post("/register", )
		// authRouter.Post("/session", )
	})
	return nil
}
