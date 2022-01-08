package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gocache/common/response"
	"gocache/common/router"
	"gocache/service"
	"gocache/utils"
	"gocache/utils/errors"
)

type CacheController interface {
	Set(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Flush(http.ResponseWriter, *http.Request)
}

type cacheController struct {
	router       router.Router
	cacheService service.CacheService
}

func NewCacheController(
	r *router.Router,
	cacheService *service.CacheService) CacheController {
	return &cacheController{
		router:       *r,
		cacheService: *cacheService,
	}
}

func (c *cacheController) Set(w http.ResponseWriter, r *http.Request) {
	var body interface{}

	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	err, httpStatusCode := c.router.Check(r)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	key := r.URL.Query().Get(utils.Key)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response.JSONResponse(nil, fmt.Errorf(errors.MissingKey)))
		return
	}

	byt, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(byt))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	if err = json.Unmarshal(byt, &body); err != nil {
		body = string(byt)
	}

	c.cacheService.Set(key, body)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.JSONResponse(body, nil))
	return
}

func (c *cacheController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	err, httpStatusCode := c.router.Check(r)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	key := r.URL.Query().Get(utils.Key)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response.JSONResponse(nil, fmt.Errorf(errors.MissingKey)))
		return
	}

	err, value := c.cacheService.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.JSONResponse(value, nil))
	return
}

func (c *cacheController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	err, httpStatusCode := c.router.Check(r)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	key := r.URL.Query().Get(utils.Key)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response.JSONResponse(nil, fmt.Errorf(errors.MissingKey)))
		return
	}

	if err = c.cacheService.Delete(key); err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.JSONResponse(true, nil))
	return
}

func (c *cacheController) Flush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utils.ContentType, utils.ApplicationJSON)
	err, httpStatusCode := c.router.Check(r)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write(response.JSONResponse(nil, err))
		return
	}

	c.cacheService.Flush()

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.JSONResponse(true, nil))
	return
}
