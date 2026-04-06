package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"password-manager/internal/api/routes"
	"time"

	"github.com/gorilla/handlers"
)

// p8my78rMkUQF0F0Bhfx68zHX4yZh1oYk

/*
	proof of concept
		- generate encrpytion key on sign up, require it for the user to sign in to access their vault
		- also require email/password

	func main() {
		encryptionKey, err := utils.GetTerminalInput("enter your 32 byte encryption key: ")
		if err != nil {
			log.Fatal(err)
		}

		encryptor, err := security.NewAESGCM([]byte(strings.TrimSpace(encryptionKey)))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("successfully read encryption key!")

		password, err := utils.GetTerminalInput("enter a password to be encrypted: ")
		if err != nil {
			log.Fatal(err)
		}

		encryptedPassword, err := encryptor.Encrypt([]byte(password))
		log.Println(encryptedPassword)
	}
*/

func main() {
	router, err := routes.InitializeRoutes()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the server
	srv := &http.Server{

		// CORS...
		Handler: handlers.CORS(
			// handlers.AllowedOrigins(config.AllowedOrigins()),
			handlers.AllowedMethods([]string{
				"GET", "POST", "HEAD", "OPTIONS", "PUT",
			}),
			handlers.AllowedHeaders([]string{
				"X-Requested-With", "Content-Type", "Authorization",
			}),
			handlers.AllowCredentials(),
		)(router),
		Addr: "127.0.0.1:8000",

		// Enforce timeouts for server
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop

		log.Println("[server] closing server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("[server] [error] closing server: %v", err)
		}
	}()

	log.Printf("[server] started at %s", "127.0.0.1:8000")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[server] [error]: %v", err)
	}

	log.Println("[server] closed successfully")
}
