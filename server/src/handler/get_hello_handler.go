package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/masakizk/go/swagger-playground/restapi/operations"
)

func GetHello(_ operations.GetHelloParams) middleware.Responder {
	payload := &operations.GetHelloOKBody{
		Message: "Hello World",
	}

	return operations.NewGetHelloOK().WithPayload(payload)
}
