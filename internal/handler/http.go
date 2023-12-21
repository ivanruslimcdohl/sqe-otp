package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type request struct {
	RqBody interface{} `json:"rq_body"`
}

func ParseRqBody(r *http.Request, rqBody interface{}) error {
	if reflect.TypeOf(rqBody).Kind() != reflect.Pointer {
		return fmt.Errorf("rqBody should be a pointer")
	}

	byBody, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %s", err.Error())
	}

	rqBodyWrapper := request{
		RqBody: rqBody,
	}
	err = json.Unmarshal(byBody, &rqBodyWrapper)
	if err != nil {
		return fmt.Errorf("failed to parse request body: %s", err.Error())
	}

	return nil
}
