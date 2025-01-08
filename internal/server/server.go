package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/goomu/internal/config"
	"github.com/osamikoyo/goomu/internal/handler"
	"github.com/osamikoyo/goomu/pkg/loger"
)

type Server struct{
	logger loger.Logger
	e *echo.Echo
	handler handler.Handler
	config config.Config
}

func New() *Server {
	return &Server{
		config: config.New(),
		logger: loger.New(),
		e: echo.New(),
		handler: handler.New(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.logger.Info().Msg("Welcome goomu!! Started preparing...")
	s.logger.Info().Msg("Register handlers")

	s.handler.RegisterRoutes(s.e)

	go func ()  {
		<- ctx.Done()
		s.e.Shutdown(ctx)
		s.logger.Info().Msg("Server shutdowned!")	
	}()

	s.logger.Info().Msg("Starting server on port 8080!")

	return s.e.Start(fmt.Sprintf("localhost:%d", s.config.Port))
}