package service

import (
	"workmate/internal/model"
	"workmate/internal/repository"
)

type URLService interface {
	GetUrlByID(id []int) (int, error)
	CheckLinksStatusByUrl(urls []string) (model.CheckLinksStatusByUrlResponse, error)
}

type Service struct {
	URLService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		URLService: NewURLService(repo.URLStorage),
	}
}
