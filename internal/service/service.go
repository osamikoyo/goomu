package service

import (
	"github.com/osamikoyo/goomu/internal/filer"
	"github.com/osamikoyo/goomu/pkg/loger"
)

type FilerService struct{
	Storage filer.Filer
	Logger loger.Logger
}

func New() *FilerService {
	return &FilerService{
		Logger: loger.New(),
		Storage: filer.New(),
	}
}
