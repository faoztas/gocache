package response

import (
	"encoding/json"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func JSONResponse(data interface{}, err error) []byte {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	byt, err := json.Marshal(Response{Data: data, Message: msg})
	if err != nil {
		return nil
	}

	return byt
}
