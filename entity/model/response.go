package model

import "strings"

type Response struct {
	Status  bool
	Message string
	Errors  interface{}
	Data    interface{}
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
}