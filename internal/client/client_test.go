package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	ast "github.com/stretchr/testify/assert"
)

func TestClient_Get(t *testing.T) {
	assert := ast.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type fields struct {
		Host    string
		Timeout time.Duration
	}

	tests := []struct {
		name     string
		fields   fields
		path     string
		wantCode int
		wantBody []byte
		wantErr  bool
	}{
		{
			name: "200 OK",
			fields: fields{
				Host:    "https://example.com",
				Timeout: 300,
			},
			path:     "posts",
			wantCode: 200,
			wantBody: []byte("mocked"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := Client{
				Host:    tt.fields.Host,
				Timeout: tt.fields.Timeout,
			}

			httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", client.Host, tt.path),
				httpmock.NewStringResponder(200, "mocked"),
			)

			gotCode, gotBody, err := client.Get(tt.path)

			assert.Equal(tt.wantCode, gotCode)
			assert.Equal(tt.wantBody, gotBody)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
