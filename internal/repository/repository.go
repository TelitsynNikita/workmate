package repository

import "workmate/internal/model"

type URLStorage interface {
	GetUrlByIDs(ids []int) ([]model.PersistentStorageData, error)
	GetLinksByUrl(urls []string) (int, []string, error)
	WriteDataToFileAndLocalStorage() error
	ReadFileToLocalStorage() error
}

type Repository struct {
	URLStorage
}

func NewRepository() *Repository {
	return &Repository{
		URLStorage: NewPersistentURLStorage(),
	}
}
