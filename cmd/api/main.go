package main

import (
	"context"
	"log"
	userHttp "mygogo/internal/user/adapters/inbound/http"
	"mygogo/internal/user/adapters/outbound/persistence"
	userApp "mygogo/internal/user/application"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/mygogo")

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	defer db.Close(context.Background())

	router := mux.NewRouter()
	userService := userApp.NewUserService(persistence.NewPostgresUserRepository(db))
	userController := userHttp.NewUserController(userService)
	userHttp.UserRouter(router, userController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running on port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
