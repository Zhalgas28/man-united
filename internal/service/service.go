package service

import (
	"man_utd/entity"
	"man_utd/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type News interface {
}

type Service struct {
	Authorization
	News
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthPostgres(repo),
		News:          nil,
	}
}
