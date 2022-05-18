package error_handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	testCase := []struct {
		name   string
		error  Error
		output string
	}{
		{
			name: "success-response",
			error: Error{
				Code:    http.StatusUnprocessableEntity,
				Message: "validation error",
			},
			output: "validation error",
		},
	}

	t.Parallel()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.error.Error()
			assert.Equal(t, output, tc.output)
		})
	}
}

func TestValidationErr(t *testing.T) {
	testCase := []struct {
		name   string
		err    error
		output Error
	}{
		{
			name: "success-response",
			err:  errors.New("validation error"),
			output: Error{
				Code:    http.StatusUnprocessableEntity,
				Message: "validation error",
			},
		},
	}

	t.Parallel()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := ValidationErr(tc.err)
			assert.Equal(t, output, tc.output)
		})
	}
}

func TestRequestErr(t *testing.T) {
	testCase := []struct {
		name   string
		err    error
		output Error
	}{
		{
			name: "success-response",
			err:  errors.New("bad request"),
			output: Error{
				Code:    http.StatusBadRequest,
				Message: "bad request",
			},
		},
	}

	t.Parallel()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := RequestErr(tc.err)
			assert.Equal(t, output, tc.output)
		})
	}
}

func TestCreateError(t *testing.T) {
	testCase := []struct {
		name   string
		code   int
		err    error
		output Error
	}{
		{
			name: "success-response",
			code: http.StatusUnprocessableEntity,
			err:  errors.New("something went wrong"),
			output: Error{
				Code:    http.StatusUnprocessableEntity,
				Message: "something went wrong",
			},
		},
	}

	t.Parallel()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := createError(tc.code, tc.err)
			assert.Equal(t, output, tc.output)
		})
	}
}
