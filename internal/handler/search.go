package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/goomu/internal/filer/models"
)

func (h Handler) searchHandler(c echo.Context) error {
	var filter models.Searchfilter
	if err := c.Bind(&filter);err != nil{
		return err
	}

	files, err := h.FileService.Search(filter)
	if err != nil{
		return err
	}

	return c.JSON(http.StatusOK, files)
}