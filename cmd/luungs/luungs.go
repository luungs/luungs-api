package luungs

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handler "github.com/luungs/luungs-api/internal/http/hanlder"
	"github.com/luungs/luungs-api/internal/repository"
	"github.com/luungs/luungs-api/internal/service"
)

type APIServer struct {
    address string
    db *sql.DB
}

func NewAPIServer (address string, db *sql.DB) *APIServer {
    return &APIServer {
        address: address,
        db: db,
    }
}

func (s *APIServer) Run() error {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(middleware.URLFormat)

    userRepository := repository.NewUserRepository(s.db)
    userService := service.NewUserService(userRepository)
    userHandler := handler.NewUserHandler(userService)

    router.Route("/api/v1", func(router chi.Router) { 
        router.Route("/users", func(router chi.Router) {
            router.Get("/", userHandler.HandleGetAllUsers)
        })
    })

    log.Print("Listening on: ", s.address)

    return http.ListenAndServe(s.address, router)
}
