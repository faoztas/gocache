package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gocache/common/response"
)

func TestRouter(t *testing.T) {
	rout := NewRouter()

	rout.GET("/get", func(writer http.ResponseWriter, request *http.Request) {})
	rout.POST("/set", func(writer http.ResponseWriter, request *http.Request) {})
	rout.DELETE("/delete", func(writer http.ResponseWriter, request *http.Request) {})
	rout.DELETE("/flush", func(writer http.ResponseWriter, request *http.Request) {})

	routeList := rout.GetMethods()

	if routeList == nil || len(*routeList) == 0 {
		t.Errorf("route list cannot be empty")
	}

	get := (*routeList)[http.MethodGet]
	if get == nil || len(get) == 0 {
		t.Errorf("get method list cannot be empty")
	}

	post := (*routeList)[http.MethodPost]
	if post == nil || len(post) == 0 {
		t.Errorf("post method list cannot be empty")
	}

	del := (*routeList)[http.MethodDelete]
	if del == nil || len(del) == 0 {
		t.Errorf("delete method list cannot be empty")
	}

	testRequest := httptest.NewRequest(http.MethodDelete, "/flush", nil)
	err, statusCode := rout.Check(testRequest)
	if err != nil {
		t.Error(err)
	}

	if statusCode > 500 {
		t.Errorf("status code problem")
	}

	testRequest = httptest.NewRequest(http.MethodPut, "/set", nil)
	err, statusCode = rout.Check(testRequest)
	if err == nil {
		t.Error("unexpected method check")
	}

	if statusCode > 500 {
		t.Errorf("status code problem")
	}

	testRequest = httptest.NewRequest(http.MethodPatch, "/set", nil)
	err, statusCode = rout.Check(testRequest)
	if err == nil {
		t.Error("unexpected method check")
	}

	if statusCode > 500 {
		t.Errorf("status code problem")
	}

	dataByte := response.JSONResponse(true, nil)
	if string(dataByte) != "{\"data\":true,\"message\":\"\"}" {
		t.Errorf("want %v, got %v", true, string(dataByte))
	}
}
