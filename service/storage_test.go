package service

import (
	"testing"

	"github.com/faoztas/gocache/repository"
	"github.com/faoztas/gocache/utils"
)

func TestStorageService_Save(t *testing.T) {
	path := utils.GenerateFilePath()
	r := repository.NewStorageRepository(&path)
	c := repository.NewCacheRepository()
	s := NewStorageService(&r, &c)

	c.Set(testData[utils.Key], testData[utils.Value])

	err := s.Save()
	if err != nil {
		t.Error(err)
	}
}

func TestStorageService_Load(t *testing.T) {
	path := utils.GenerateFilePath()
	r := repository.NewStorageRepository(&path)
	c := repository.NewCacheRepository()
	s := NewStorageService(&r, &c)

	err := s.Load()
	if err != nil {
		t.Error(err)
	}

	err, value := c.Get(testData[utils.Key])
	if err != nil {
		t.Error(err)
	}

	if value != testData[utils.Value] {
		t.Errorf("want %v, got %v", value, testData[utils.Value])
	}
}
