package repository

import (
	"fmt"
	"sync"

	"gocache/utils/errors"
)

type CacheRepository interface {
	Set(string, interface{})
	Get(string) (error, interface{})
	Delete(string) error
	Flush()
	GetAll() map[string]interface{}
	Load(map[string]interface{})
}

type cacheRepository struct {
	sync.RWMutex
	cache map[string]interface{}
}

func NewCacheRepository() CacheRepository {
	return &cacheRepository{
		cache: make(map[string]interface{}),
	}
}

func (r *cacheRepository) Set(key string, value interface{}) {
	r.Lock()
	defer r.Unlock()

	r.cache[key] = value
}

func (r *cacheRepository) Get(key string) (error, interface{}) {
	r.Lock()
	defer r.Unlock()

	value := r.cache[key]
	if value == nil {
		return fmt.Errorf(errors.CacheRecordNotFound), nil
	}
	return nil, value
}

func (r *cacheRepository) Delete(key string) error {
	r.Lock()
	defer r.Unlock()

	if r.cache[key] == nil {
		return fmt.Errorf(errors.CacheRecordNotFound)
	}
	delete(r.cache, key)

	return nil
}

func (r *cacheRepository) Flush() {
	r.Lock()
	defer r.Unlock()

	r.cache = make(map[string]interface{})
}

func (r *cacheRepository) GetAll() map[string]interface{} {
	r.Lock()
	defer r.Unlock()

	return r.cache
}

func (r *cacheRepository) Load(data map[string]interface{}) {
	r.Lock()
	defer r.Unlock()

	r.cache = data
}
