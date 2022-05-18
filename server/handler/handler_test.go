package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ServiceStructure/utils/error-handler"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequest(t *testing.T) {
	testCase := []struct {
		name               string
		mockRequestHandler func(ctx context.Context, req *http.Request) Response
		req                *http.Request
		actualResponse     *httptest.ResponseRecorder
		expectedResponse   *httptest.ResponseRecorder
		requestHandler     func(ctx context.Context, req *http.Request) http.HandlerFunc
	}{
		{
			name: "make rest call",
			req: &http.Request{
				Body: http.NoBody,
			},
			mockRequestHandler: func(ctx context.Context, req *http.Request) Response {
				return Response{
					Code: http.StatusOK,
					Data: "dummy-response",
				}
			},
			actualResponse:   httptest.NewRecorder(),
			expectedResponse: httptest.NewRecorder(),
			requestHandler: func(ctx context.Context, req *http.Request) http.HandlerFunc {
				return func(w http.ResponseWriter, req *http.Request) {
					response := Response{
						Code: http.StatusOK,
						Data: "dummy-response",
					}
					JSON, _ := json.Marshal(response.Data)
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write(JSON)
				}
			},
		},
	}

	ctx := context.Background()
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			Make(tc.mockRequestHandler)(tc.actualResponse, tc.req)
			tc.requestHandler(ctx, tc.req)(tc.expectedResponse, tc.req)
			assert.Equal(t, tc.actualResponse.Code, tc.expectedResponse.Code)
			assert.Equal(t, tc.actualResponse.Body, tc.expectedResponse.Body)
		})
	}
}

func TestWriteErr(t *testing.T) {
	testCase := []struct {
		name       string
		ctx        context.Context
		req        *http.Request
		w          *httptest.ResponseRecorder
		val        interface{}
		outputCode int
		outputBody *bytes.Buffer
	}{
		{
			name: "success-response",
			ctx:  context.Background(),
			req: &http.Request{
				Body: http.NoBody,
			},
			w: httptest.NewRecorder(),
			val: error_handler.Error{
				Code:    http.StatusServiceUnavailable,
				Message: "dummy-error",
			},
			outputCode: http.StatusServiceUnavailable,
			outputBody: bytes.NewBuffer([]byte(`{"code":503,"message":"dummy-error"}`)),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			writeErr(tc.ctx, tc.req, tc.w, tc.val)
			assert.Equal(t, tc.w.Code, tc.outputCode)
			assert.Equal(t, tc.w.Body, tc.outputBody)
		})
	}
}

func TestWriteToResponseWrite(t *testing.T) {
	testCase := []struct {
		name       string
		ctx        context.Context
		req        *http.Request
		w          *httptest.ResponseRecorder
		response   Response
		outputCode int
		outputBody *bytes.Buffer
	}{
		{
			name: "success-response",
			ctx:  context.Background(),
			req: &http.Request{
				Body: http.NoBody,
			},
			w: httptest.NewRecorder(),
			response: Response{
				Code: http.StatusOK,
				Data: "dummy-response",
			},
			outputCode: http.StatusOK,
			outputBody: bytes.NewBuffer([]byte("\"dummy-response\"")),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			writeToResponseWrite(tc.ctx, tc.req, tc.w, tc.response)
			assert.Equal(t, tc.w.Code, tc.outputCode)
			assert.Equal(t, tc.w.Body, tc.outputBody)
		})
	}
}

func TestSetupResponse(t *testing.T) {
	testCase := []struct {
		name   string
		w      *httptest.ResponseRecorder
		output http.Header
	}{
		{
			name: "success-response",
			w:    httptest.NewRecorder(),
			output: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			setupResponse(tc.w)
			assert.Equal(t, tc.w.Header(), tc.output)
		})
	}
}
