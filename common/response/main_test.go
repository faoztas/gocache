package response

import (
	"encoding/json"
	"testing"
)

func TestJSONResponse(t *testing.T) {
	var r Response
	var data = true

	byt := JSONResponse(data, nil)
	if err := json.Unmarshal(byt, &r); err != nil {
		t.Error(err)
	}

	if r.Data != data {
		t.Errorf("want %v, got %v", r.Data, data)
	}
}
