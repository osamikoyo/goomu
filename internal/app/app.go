package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/osamikoyo/goomu/internal/server"
)

func App()  {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server.New().Run(ctx)
}