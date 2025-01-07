package models

import "mime/multipart"

type File struct{
	Header *multipart.FileHeader
	Group string
}

type Searchfilter struct{
	Ext string `json:"ext"`
	Name string `json:"name"` //without ext
	Type string `json:"group"`
}