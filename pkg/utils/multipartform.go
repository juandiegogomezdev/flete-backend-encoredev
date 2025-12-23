package utils

import (
	"fmt"
	"log"
	"net/http"
)

const (
	KB int64 = 1 << 10
	MB int64 = 1 << 20
)

func ParseMultipartForm(r *http.Request) error {
	err := r.ParseMultipartForm(32 * KB)
	if err != nil {
		log.Println("Error reading multipartForm (utils): ", err)
		return fmt.Errorf("Error al analizar los datos del formularoi")
	}
	return nil
}
