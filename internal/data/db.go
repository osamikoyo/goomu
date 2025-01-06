package data

type FileSaver struct{}

func New() *FileSaver {
	return new(FileSaver)
}


func (f *FileSaver) Save()