package app

import "github.com/osamikoyo/goomu/internal/server"

func App()  {
	server.New().Run()
}