package ping

import (
	"context"
	"testing"

	"ServiceStructure/server/handler"

	"github.com/stretchr/testify/assert"
)

func TestPingService(t *testing.T) {
	testCase := []struct {
		name       string
		ctx        context.Context
		mockESPing func(ctx context.Context) error
		output     handler.Response
	}{
		{
			name: `ping`,
			ctx:  context.Background(),
			mockESPing: func(ctx context.Context) error {
				return nil
			},
			output: handler.Response(struct {
				Code int
				Data interface{}
			}{
				Code: 200,
				Data: "Success",
			}),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			output := Ping(tc.ctx, nil)
			assert.Equal(t, tc.output, output)
		})
	}
}
