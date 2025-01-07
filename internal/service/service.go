package service

import (
	"github.com/osamikoyo/goomu/internal/filer"
	"github.com/osamikoyo/goomu/internal/filer/models"
)

type Service interface{
	Save(models.File) error
	Search(filter models.Searchfilter) ([]models.File, error)
	Display(group_name string) ([]models.File, error)
}

type FilerService struct{
	Storage filer.Filer
}

