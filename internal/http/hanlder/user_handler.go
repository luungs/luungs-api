package handler

import (
	"net/http"

	"github.com/luungs/luungs-api/internal/service"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service} 
}

func (h *UserHandler) HandleGetAllUsers (w http.ResponseWriter, r *http.Request) {
}
