package repository

import (
	"encoding/json"
	"testing"

	"gocache/utils"
)

func TestStorageRepository_SetGet(t *testing.T) {
	var mdl map[string]interface{}
	path := utils.GenerateFilePath()
	r := NewStorageRepository(&path)

	data, err := json.Marshal(map[string]interface{}{testData[utils.Key]: testData[utils.Value]})
	if err != nil {
		t.Error(err)
	}

	err = r.Set(data)
	if err != nil {
		t.Error(err)
	}

	err, data = r.Get()
	if err != nil {
		t.Error(err)
	}

	if err = json.Unmarshal(data, &mdl); err != nil {
		t.Error(err)
	}

	if mdl[testData[utils.Key]] != testData[utils.Value] {
		t.Errorf("want %v, got %v", mdl[testData[utils.Key]], testData[utils.Value])
	}
}
