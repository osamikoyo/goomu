package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) get_AllHandler(c echo.Context) error {
	group := c.Param("group")

	files, err := h.FileService.Get_All(group)
	if err != nil{
		return err
	}

	return c.JSON(http.StatusOK, files)
}