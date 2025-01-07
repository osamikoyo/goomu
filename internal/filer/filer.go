package filer

import "github.com/osamikoyo/goomu/internal/filer/models"

type Filer interface{
	Save(file models.File) error
	Display() ([]models.File, error)
	Search(name string) (models.File, error)
}

type GoomuFiler struct{

}

