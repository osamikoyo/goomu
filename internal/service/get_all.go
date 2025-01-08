package service

import "github.com/osamikoyo/goomu/internal/filer/models"

func (f *FilerService) Get_All(group string) ([]models.GittingFile, error){
	var files []models.GittingFile

	fils, err := f.Storage.Get_All()
	if err != nil{
		return []models.GittingFile{}, nil
	}

	if group == "" {
		return fils, nil		
	} else {
		for _, f := range fils{
			if f.Group == group {
				files = append(files, f)
			}
		}
	
		return files, nil
	}
}
	
