package service

import (
	"errors"
	"path/filepath"

	"github.com/osamikoyo/goomu/internal/filer/models"
)

func (f *FilerService) Search(filter models.Searchfilter) ([]models.GittingFile, error){
	var files []models.GittingFile

	fls, err := f.Storage.Get_All()
	if err != nil{
		return files, err
	}

	if filter.Ext == "" && filter.Name == "" {
		return files, errors.New("to little arguments")
	} else if filter.Ext == "" {
		for _, fl :=  range fls{
			if fl.Name == filter.Name {
				files = append(files, fl)
			}
		}	
	} else if filter.Name != "" && filter.Ext != ""{
		for _, fl :=  range fls{
			if filepath.Ext(fl.Path) == filter.Ext && fl.Name == filter.Name{
				files = append(files, fl)
			}
		}
	} else {
		for _, fl :=  range fls{
			if filepath.Ext(fl.Path) == filter.Ext {
				files = append(files, fl)
			}
		}	

	}

	return files, nil
}