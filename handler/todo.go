package handler

import (
	"net/http"

	"github.com/pdrum/todo/request"
	"github.com/pdrum/todo/service"

	"github.com/labstack/echo"
)

type ItemHandler struct {
	Service service.ItemService
}

func (i ItemHandler) Create(c echo.Context) error {
	req := request.CreateItemRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp, err := i.Service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i ItemHandler) ListAll(c echo.Context) error {
	resp, err := i.Service.ListAll()
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
