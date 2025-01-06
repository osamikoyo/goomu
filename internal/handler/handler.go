package handler

import "net/http"

type Handler struct{
	
}

func s(w http.ResponseWriter, r *http.Request){
	d, f, err := r.FormFile("file")
}