package handler

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yukikamome316/httpmock-test/internal/api"
	"github.com/yukikamome316/httpmock-test/internal/client"
)

type Handler struct {
	Gw client.ApiGateway
}

func NewHandler(gw client.ApiGateway) Handler {
	return Handler{
		Gw: gw,
	}
}

// ServiceをHandlerに持たせるべきなんだろうけど省略
func (h Handler) GetPosts(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("str converting failed: %s", idStr)
	}
	return api.GetPostsApi(h.Gw, id)
}
