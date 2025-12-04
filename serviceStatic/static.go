package serviceStatic

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var dist embed.FS

//encore:service
type ServiceStatic struct {
	staticHandler http.Handler
}

func initServiceStatic() (*ServiceStatic, error) {
	var assets, err = fs.Sub(dist, "dist")

	if err != nil {
		return nil, err
	}

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.FS(assets)))
	return &ServiceStatic{staticHandler: staticHandler}, nil
}

//encore:api public raw path=/static/*path tag:static
func (s *ServiceStatic) ServeStaticFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Using endpotin /static/*path")
	s.staticHandler.ServeHTTP(w, req)
}
