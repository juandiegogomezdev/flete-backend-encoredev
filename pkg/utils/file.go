package utils

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

func ExtractFileFromMultipartform(r *http.Request, formFilename string, maxSize int64) (multipart.File, *multipart.FileHeader, error) {

	file, header, err := r.FormFile(formFilename)

	if err != nil {
		log.Println("Error reading file: ", err)
		return nil, nil, fmt.Errorf("Error al leer el archivo")
	}

	if header.Size > maxSize {
		log.Println("too large file")
		file.Close()
		return nil, nil, fmt.Errorf("El archivo es muy pesado")
	}

	return file, header, nil

}
