package utils

import (
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"encore.dev/beta/errs"
)

func GetExtensionFromFilename(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}

func ExtractFileFromMultipartform(r *http.Request, formFilename string, maxSize int64) (multipart.File, *multipart.FileHeader, error) {

	file, header, err := r.FormFile(formFilename)

	if err != nil {
		log.Println("Error reading file: ", err)
		return nil, nil, &errs.Error{
			Code:    errs.Internal,
			Message: "Error al leer el archivo",
		}
	}

	if header.Size > maxSize {
		log.Println("too large file")
		file.Close()
		return nil, nil, &errs.Error{
			Code:    errs.Internal,
			Message: "El archivo es muy pesado",
		}
	}

	return file, header, nil

}
