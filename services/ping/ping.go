package ping

import (
	"context"
	"log"
	"net/http"

	"ServiceStructure/server/handler"
)

func Ping(ctx context.Context, _ *http.Request) handler.Response {
	log.Printf("Health check")

	return handler.Response{
		Code: http.StatusOK,
		Data: "Success",
	}
}
