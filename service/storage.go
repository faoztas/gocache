package service

import (
	"encoding/json"

	"github.com/faoztas/gocache/repository"
)

type StorageService interface {
	Save() error
	Load() error
}

type storageService struct {
	storageRepository repository.StorageRepository
	cacheRepository   repository.CacheRepository
}

func NewStorageService(
	storageRepository *repository.StorageRepository,
	cacheRepository *repository.CacheRepository) StorageService {
	return &storageService{
		storageRepository: *storageRepository,
		cacheRepository:   *cacheRepository,
	}
}

func (s *storageService) Save() error {
	data, err := json.Marshal(s.cacheRepository.GetAll())
	if err != nil {
		return err
	}

	return s.storageRepository.Set(data)
}

func (s *storageService) Load() error {
	var mdl map[string]interface{}

	err, data := s.storageRepository.Get()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &mdl); err != nil {
		return err
	}

	s.cacheRepository.Load(mdl)
	return nil
}
