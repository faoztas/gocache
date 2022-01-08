package repository

import (
	"testing"

	"gocache/utils"
)

var testData = map[string]string{
	utils.Key:   "hello",
	utils.Value: "world!",
}

func TestCacheRepository_SetGet(t *testing.T) {
	r := NewCacheRepository()

	r.Set(testData[utils.Key], testData[utils.Value])

	err, value := r.Get(testData[utils.Key])
	if err != nil {
		t.Error(err)
	}

	if value != testData[utils.Value] {
		t.Errorf("want %v, got %v", value, testData[utils.Value])
	}
}

func TestCacheRepository_Delete(t *testing.T) {
	r := NewCacheRepository()

	r.Set(testData[utils.Key], testData[utils.Value])

	err := r.Delete(testData[utils.Key])
	if err != nil {
		t.Error(err)
	}

	err, _ = r.Get(testData[utils.Key])
	if err == nil {
		t.Error("delete operation failed")
	}
}

func TestCacheRepository_Flush(t *testing.T) {
	r := NewCacheRepository()

	r.Set(testData[utils.Key], testData[utils.Value])

	r.Flush()

	err, _ := r.Get(testData[utils.Key])
	if err == nil {
		t.Error("flush operation failed")
	}
}

func TestCacheRepository_GetAll(t *testing.T) {
	r := NewCacheRepository()

	r.Set(testData[utils.Key], testData[utils.Value])

	if r.GetAll() == nil && len(r.GetAll()) == 0 {
		t.Errorf("cache cannot be empty")
	}
}

func TestCacheRepository_Load(t *testing.T) {
	r := NewCacheRepository()

	r.Load(map[string]interface{}{testData[utils.Key]: testData[utils.Value]})

	err, value := r.Get(testData[utils.Key])
	if err != nil {
		t.Error(err)
	}

	if value != testData[utils.Value] {
		t.Errorf("want %v, got %v", value, testData[utils.Value])
	}
}
