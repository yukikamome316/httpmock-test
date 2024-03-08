package api

import (
	"fmt"
	"net/http"

	"github.com/yukikamome316/httpmock-test/internal/client"
)

const postsPath = "posts/%d"

func GetPostsApi(gw client.ApiGateway, id int) error {
	params := make(map[string]interface{})
	params["post_id"] = id

	path := fmt.Sprintf(postsPath, id)
	code, body, err := gw.Get(path)
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf("(%d) GET %s: %s", code, path, body)
	}

	return nil
}
