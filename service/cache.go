package service

import (
	"github.com/faoztas/gocache/repository"
)

type CacheService interface {
	Set(string, interface{})
	Get(string) (error, interface{})
	Delete(string) error
	Flush()
}

type cacheService struct {
	cacheRepository repository.CacheRepository
}

func NewCacheService(
	cacheRepository *repository.CacheRepository) CacheService {
	return &cacheService{
		cacheRepository: *cacheRepository,
	}
}

func (s *cacheService) Set(key string, value interface{}) {
	s.cacheRepository.Set(key, value)
}

func (s *cacheService) Get(key string) (error, interface{}) {
	return s.cacheRepository.Get(key)
}

func (s *cacheService) Delete(key string) error {
	return s.cacheRepository.Delete(key)
}

func (s *cacheService) Flush() {
	s.cacheRepository.Flush()
}
