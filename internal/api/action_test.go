package api

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	ast "github.com/stretchr/testify/assert"
	"github.com/yukikamome316/httpmock-test/internal/client"
)

const privateHost = "https://example.com"

func TestGetPostsApi(t *testing.T) {
	assert := ast.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type args struct {
		gw client.ApiGateway
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "200 OK",
			args: args{
				gw: &client.Client{
					Host:    privateHost,
					Timeout: 300,
				},
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gw := tt.args.gw
			id := tt.args.id
			httpmock.RegisterResponder("GET", fmt.Sprintf("%s/posts/%d", privateHost, id),
				httpmock.NewStringResponder(200, "mocked"),
			)

			err := GetPostsApi(gw, id)

			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
