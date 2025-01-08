package service

import (
	"github.com/osamikoyo/goomu/internal/filer/models"
)

func (f *FilerService) Save(file models.File) error {
	return f.Storage.Save(file)
}