package repository

import (
	"os"
	"sync"

	"gocache/utils"
)

type StorageRepository interface {
	Set([]byte) error
	Get() (error, []byte)
}

type storageRepository struct {
	sync.RWMutex
	path string
}

func NewStorageRepository(
	path *string) StorageRepository {
	return &storageRepository{
		path: *path,
	}
}

func (r *storageRepository) Set(data []byte) error {
	r.Lock()
	defer r.Unlock()

	if err := os.WriteFile(r.path, data, utils.Perm); err != nil {
		return err
	}

	return nil
}

func (r *storageRepository) Get() (error, []byte) {
	r.Lock()
	defer r.Unlock()

	data, err := os.ReadFile(r.path)
	if err != nil {
		return err, nil
	}

	return nil, data
}
