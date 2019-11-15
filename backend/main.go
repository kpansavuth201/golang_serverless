package main

import (
	"log"
	"net/http"
	"path"
	"strings"
)

const (
	port      = "8080"
	publicDir = "./public"
)

func Handler(publicDir string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}

		// the all 404 gonna be served as root
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}

func main() {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: Handler(publicDir),
	}
	log.Printf("Listening on port %s", port)
	err := server.ListenAndServe()
	panic(err)
}
