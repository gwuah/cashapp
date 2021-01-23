package core

import (
	"log"
	"net/http"

	"github.com/rs/xid"
)

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func GenerateRef() string {
	return xid.New().String()
}

func Error(err error, m *string) Response {
	log.Println(err)

	var message string
	if m == nil {
		message = "request failed"
	} else {
		message = StringValue(m)
	}

	return Response{
		Error: true,
		Code:  http.StatusInternalServerError,
		Meta: Meta{
			Data:    nil,
			Message: message,
		},
	}
}

func Success(data *map[string]interface{}, m *string) Response {

	var message string
	if m == nil {
		message = "request successful"
	} else {
		message = StringValue(m)
	}

	return Response{
		Error: false,
		Code:  http.StatusOK,
		Meta: Meta{
			Data:    data,
			Message: message,
		},
	}
}
