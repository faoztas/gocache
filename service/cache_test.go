package service

import (
	"testing"

	"github.com/faoztas/gocache/repository"
	"github.com/faoztas/gocache/utils"
)

var testData = map[string]string{
	utils.Key:   "hello",
	utils.Value: "world!",
}

func TestCacheService_SetGet(t *testing.T) {
	r := repository.NewCacheRepository()
	s := NewCacheService(&r)

	s.Set(testData[utils.Key], testData[utils.Value])

	err, value := s.Get(testData[utils.Key])
	if err != nil {
		t.Error(err)
	}

	if value != testData[utils.Value] {
		t.Errorf("want %v, got %v", value, testData[utils.Value])
	}
}

func TestCacheService_Delete(t *testing.T) {
	r := repository.NewCacheRepository()
	s := NewCacheService(&r)

	s.Set(testData[utils.Key], testData[utils.Value])

	if err := s.Delete(testData[utils.Key]); err != nil {
		t.Error(err)
	}

	err, _ := s.Get(testData[utils.Key])
	if err == nil {
		t.Error("delete operation failed")
	}
}

func TestCacheService_Flush(t *testing.T) {
	r := repository.NewCacheRepository()
	s := NewCacheService(&r)

	s.Set(testData[utils.Key], testData[utils.Value])

	s.Flush()

	err, _ := s.Get(testData[utils.Key])
	if err == nil {
		t.Error("delete operation failed")
	}
}
