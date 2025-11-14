package service

import (
	"fmt"
	"workmate/internal/model"
	"workmate/internal/repository"
)

type URL struct {
	URLStorage repository.URLStorage
}

func NewURLService(repo repository.URLStorage) *URL {
	return &URL{
		URLStorage: repo,
	}
}

func (u *URL) GetUrlByID(ids []int) (int, error) {
	links, err := u.URLStorage.GetUrlByIDs(ids)
	if err != nil {
		return 0, err
	}

	fmt.Println(links)

	return 0, nil
}

func (u *URL) CheckLinksStatusByUrl(urls []string) (model.CheckLinksStatusByUrlResponse, error) {
	id, links, err := u.URLStorage.GetLinksByUrl(urls)
	if err != nil {
		return model.CheckLinksStatusByUrlResponse{}, err
	}

	fmt.Println(id, links)

	return model.CheckLinksStatusByUrlResponse{}, nil
}
