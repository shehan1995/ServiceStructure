package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"ServiceStructure/utils/error-handler"
)

// Response is written to http.ResponseWriter
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type RequestHandler func(ctx context.Context, req *http.Request) Response

// Make creates a http handler from a request handler func
func Make(fn RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		defer func(ctx context.Context) {
			if val, ok := RecoverFromPanic(ctx, recover()); ok {
				writeErr(ctx, req, w, val)
			}
		}(ctx)
		setupResponse(w)
		res := fn(ctx, req)
		writeToResponseWrite(ctx, req, w, res)
	}
}

// writeErr write error to response
func writeErr(ctx context.Context, req *http.Request, w http.ResponseWriter, val interface{}) {
	response := Response{
		Code: http.StatusServiceUnavailable,
		Data: val,
	}
	setupResponse(w)
	writeToResponseWrite(ctx, req, w, response)
}

// writeToResponseWrite write response
func writeToResponseWrite(ctx context.Context, req *http.Request, w http.ResponseWriter, response Response) {
	log.Printf("Http response is : %v", response)
	JSON, err := json.Marshal(response.Data)
	if err != nil {
		log.Printf("json marshal failed %v", err)
	}
	w.WriteHeader(response.Code)
	_, err = w.Write(JSON)
	if err != nil {
		log.Printf("json write failed %v", err)
	}
	err = req.Body.Close()
	if err != nil {
		log.Printf("error while closing the body %v", err)
	}
}

// setupResponse setups response for api content-type and can also set cors policy headers
func setupResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// RecoverFromPanic used to recover from panic and log the error
func RecoverFromPanic(ctx context.Context, err interface{}) (interface{}, bool) {
	if err != nil {
		log.Printf("Panic Recovered!: %v \n %v", err, string(debug.Stack()))
		return error_handler.ServerError, true
	}
	return nil, false
}
