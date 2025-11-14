package repository

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"workmate/internal/model"
)

type PersistentURLStorage struct {
	localStorageMutex      sync.RWMutex
	localStorage           map[int]string
	persistentStorageMutex sync.RWMutex
}

func NewPersistentURLStorage() *PersistentURLStorage {
	return &PersistentURLStorage{
		localStorage: make(map[int]string),
	}
}

func (p *PersistentURLStorage) GetLinksByUrl(urls []string) (int, []string, error) {
	if len(p.localStorage) == 0 {
		err := p.ReadFileToLocalStorage()
		if err != nil {
			return 0, nil, err
		}
	}

	sort.Strings(urls)
	urlsJoined := strings.Join(urls, ",")
	encodeUrlStrings := base64.StdEncoding.EncodeToString([]byte(urlsJoined))

	var theMostIndexInMap = 0
	for key, value := range p.localStorage {
		if encodeUrlStrings == value {
			return key, urls, nil
		}

		if theMostIndexInMap < key {
			theMostIndexInMap = key
		}
	}

	p.localStorageMutex.Lock()
	p.localStorage[theMostIndexInMap+1] = encodeUrlStrings
	p.localStorageMutex.Unlock()

	return theMostIndexInMap + 1, urls, nil
}

func (p *PersistentURLStorage) GetUrlByIDs(ids []int) ([]model.PersistentStorageData, error) {
	if len(p.localStorage) == 0 {
		err := p.ReadFileToLocalStorage()
		if err != nil {
			return nil, err
		}
	}

	var links []model.PersistentStorageData
	for _, id := range ids {
		p.localStorageMutex.RLock()
		encodeString, ok := p.localStorage[id]
		if !ok {
			return nil, fmt.Errorf("there is no data by id: %d", id)
		}
		p.localStorageMutex.RUnlock()

		data, err := base64.StdEncoding.DecodeString(encodeString)
		if err != nil {
			return nil, err
		}

		links = append(links, model.PersistentStorageData{
			ID:          id,
			LinkedLinks: strings.Split(string(data), ","),
		})
	}

	if len(links) == 0 {
		return nil, fmt.Errorf("there is no data by ids: %v", ids)
	}

	return links, nil
}

func (p *PersistentURLStorage) ReadFileToLocalStorage() error {
	p.persistentStorageMutex.RLock()
	file, err := os.ReadFile("persistent_storage.txt")
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = p.CreatePersistentStorage()
		if err != nil {
			return err
		}

		file, err = os.ReadFile("persistent_storage.txt")
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	var data map[int]string

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}
	p.persistentStorageMutex.RUnlock()

	p.localStorageMutex.Lock()
	p.localStorage = data
	p.localStorageMutex.Unlock()

	return nil
}

func (p *PersistentURLStorage) WriteDataToFileAndLocalStorage() error {
	p.persistentStorageMutex.Lock()
	defer p.persistentStorageMutex.Unlock()
	err := exec.Command("/bin/bash", "-c", "echo > ./persistent_storage.txt").Run()
	if err != nil {
		return err
	}

	data, err := json.Marshal(p.localStorage)
	if err != nil {
		return err
	}

	err = os.WriteFile("persistent_storage.txt", data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (p *PersistentURLStorage) CreatePersistentStorage() error {
	file, err := os.Create("persistent_storage.txt")
	if err != nil || file == nil {
		return err
	}
	defer file.Close()

	return nil
}
