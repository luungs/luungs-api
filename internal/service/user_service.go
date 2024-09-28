package service

import "github.com/luungs/luungs-api/internal/repository"

type UserService struct {
    r *repository.UserRepository
}

func NewUserService (r *repository.UserRepository) *UserService {
    return &UserService{r: r}
}
