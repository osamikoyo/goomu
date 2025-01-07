package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/goomu/internal/filer"
	"github.com/osamikoyo/goomu/pkg/loger"
)

type Handler struct{
	Filer filer.Filer
	SaveDir string
}


func errorRoute(h echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		if err := h(c);err != nil{
			loger.New().Error().Err(err)
		}
		return nil
	}
}

func (h Handler) RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.POST("/file/upload", errorRoute(h.Save))
}
