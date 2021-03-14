package handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/masakizk/go/swagger-playground/restapi/operations"
)

func PostHello(p operations.PostHelloParams) middleware.Responder {
	payload := &operations.PostHelloOKBody{
		Message: fmt.Sprintf("Hello %s", *p.Name),
	}
	return operations.NewPostHelloOK().WithPayload(payload)
}
