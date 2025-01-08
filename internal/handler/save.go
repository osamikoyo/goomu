package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/goomu/internal/filer/models"
)

func (h Handler) saveHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil{
		return err
	}

	return h.FileService.Save(models.File{
		Header: file,
		Group: c.FormValue("group"),
	})
}