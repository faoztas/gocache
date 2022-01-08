package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gocache/common/response"
	"gocache/common/router"
	"gocache/repository"
	"gocache/service"
	"gocache/utils"
)

var testData = map[string]string{
	utils.Key:   "hello",
	utils.Value: "world!",
}

func TestCacheController_Set(t *testing.T) {
	var r response.Response

	rout := router.NewRouter()
	repo := repository.NewCacheRepository()
	serv := service.NewCacheService(&repo)
	cont := NewCacheController(&rout, &serv)

	prefix := "/set"
	rout.POST(prefix, cont.Set)

	request := httptest.NewRequest(http.MethodPost, prefix, bytes.NewBufferString(testData[utils.Value]))
	q := url.Values{}
	q.Add(utils.Key, testData[utils.Key])
	request.URL.RawQuery = q.Encode()

	rWriter := httptest.NewRecorder()
	cont.Set(rWriter, request)
	resp := rWriter.Result()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		t.Errorf("status code error %d", resp.StatusCode)
	case http.StatusOK:
		break
	default:
		t.Errorf("unexpected status code error %d", resp.StatusCode)
	}

	byt, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(byt, &r); err != nil {
		t.Error(err)
	}

	if r.Data != testData[utils.Value] {
		t.Errorf("want %v, got %v", r.Data, testData[utils.Value])
	}
}

func TestCacheController_Get(t *testing.T) {
	var r response.Response

	rout := router.NewRouter()
	repo := repository.NewCacheRepository()
	serv := service.NewCacheService(&repo)
	cont := NewCacheController(&rout, &serv)

	repo.Set(testData[utils.Key], testData[utils.Value])

	prefix := "/get"
	rout.GET(prefix, cont.Get)

	request := httptest.NewRequest(http.MethodGet, prefix, nil)
	q := url.Values{}
	q.Add(utils.Key, testData[utils.Key])
	request.URL.RawQuery = q.Encode()

	rWriter := httptest.NewRecorder()
	cont.Get(rWriter, request)
	resp := rWriter.Result()

	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		t.Errorf("status code error %d", resp.StatusCode)
	case http.StatusOK:
		break
	default:
		t.Errorf("unexpected status code error %d", resp.StatusCode)
	}

	byt, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(byt, &r); err != nil {
		t.Error(err)
	}

	if r.Data != testData[utils.Value] {
		t.Errorf("want %v, got %v", r.Data, testData[utils.Value])
	}
}

func TestCacheController_Delete(t *testing.T) {
	var r response.Response

	rout := router.NewRouter()
	repo := repository.NewCacheRepository()
	serv := service.NewCacheService(&repo)
	cont := NewCacheController(&rout, &serv)

	repo.Set(testData[utils.Key], testData[utils.Value])

	prefix := "/delete"
	rout.DELETE(prefix, cont.Delete)

	request := httptest.NewRequest(http.MethodDelete, prefix, nil)
	q := url.Values{}
	q.Add(utils.Key, testData[utils.Key])
	request.URL.RawQuery = q.Encode()

	rWriter := httptest.NewRecorder()
	cont.Delete(rWriter, request)
	resp := rWriter.Result()

	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		t.Errorf("status code error %d", resp.StatusCode)
	case http.StatusOK:
		break
	default:
		t.Errorf("unexpected status code error %d", resp.StatusCode)
	}

	byt, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(byt, &r); err != nil {
		t.Error(err)
	}

	if r.Data != true {
		t.Errorf("want %v, got %v", r.Data, true)
	}
}

func TestCacheController_Flush(t *testing.T) {
	var r response.Response

	rout := router.NewRouter()
	repo := repository.NewCacheRepository()
	serv := service.NewCacheService(&repo)
	cont := NewCacheController(&rout, &serv)

	repo.Set(testData[utils.Key], testData[utils.Value])

	prefix := "/flush"
	rout.DELETE(prefix, cont.Delete)

	request := httptest.NewRequest(http.MethodDelete, prefix, nil)
	q := url.Values{}
	q.Add(utils.Key, testData[utils.Key])
	request.URL.RawQuery = q.Encode()

	rWriter := httptest.NewRecorder()
	cont.Flush(rWriter, request)
	resp := rWriter.Result()

	switch resp.StatusCode {
	case http.StatusOK:
		break
	default:
		t.Errorf("unexpected status code error %d", resp.StatusCode)
	}

	byt, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(byt, &r); err != nil {
		t.Error(err)
	}

	if r.Data != true {
		t.Errorf("want %v, got %v", r.Data, true)
	}
}
