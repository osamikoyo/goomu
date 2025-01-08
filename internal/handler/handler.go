package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/goomu/internal/config"
	"github.com/osamikoyo/goomu/internal/filer"
	"github.com/osamikoyo/goomu/internal/service"
	"github.com/osamikoyo/goomu/pkg/loger"
)

type Handler struct{
	FileService service.FilerService
	SaveDir string
	logger loger.Logger
}

func New() Handler {
	config := config.New()

	return Handler{
		FileService: service.FilerService{
			Storage: filer.New(),
		},
		SaveDir: config.Upload_Dir,
		logger: loger.New(),
	}
}

func errorRoute(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h(c)
		if err != nil {
			loger.New().Error().
				Str("path", c.Path()).
				Str("method", c.Request().Method).
				Err(err).
				Msg("Request failed")
			
			return err
		}
		return nil
	}
}

func (h Handler) RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/file/get/:group", errorRoute(h.get_AllHandler))
	e.POST("/file/get", errorRoute(h.searchHandler))
	e.POST("/file/upload", errorRoute(h.saveHandler))
}
