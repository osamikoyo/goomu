package filer

import (
	"io"
	"os"
	"path/filepath"

	"github.com/osamikoyo/goomu/internal/filer/models"
)

type Filer interface{
	Save(file models.File) error
	Get_All() ([]models.GittingFile, error)
}

type GoomuFiler struct{
	UploadDir string
}

func New() *GoomuFiler {
	return &GoomuFiler{
		UploadDir: "storage",
	}
}

func (g *GoomuFiler) Save(file models.File) error {
	uploadPath := filepath.Join(g.UploadDir, file.Group)
	_ = os.MkdirAll(uploadPath, 0755)

	dest, err := os.Create(filepath.Join(uploadPath, file.Header.Filename))
	if err != nil {
		return err
	}
	defer dest.Close()

	src, err := file.Header.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	return nil
}

func (g *GoomuFiler) Get_All() ([]models.GittingFile, error) {
	var files []models.GittingFile
	
	err := filepath.Walk(g.UploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if info.IsDir() {
			return nil
		}
		
		// Get relative path from upload directory
		relPath, err := filepath.Rel(g.UploadDir, path)
		if err != nil {
			return err
		}
		
		// Extract group (first directory in path)
		group := filepath.Dir(relPath)
		if group == "." {
			group = "" // Root directory
		}
		
		file := models.GittingFile{
			Header: nil, // We don't have FileHeader for existing files
			Name:   info.Name(),
			Size:   info.Size(),
			Group:  group,
			Path:   path,
		}
		
		files = append(files, file)
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return files, nil
}