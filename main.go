package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/axzilla/goilerplate-quickstart/assets"
	"github.com/axzilla/goilerplate-quickstart/ui/pages"
	"github.com/joho/godotenv"
)

func main() {
	InitDotEnv()
	mux := http.NewServeMux()
	SetupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(pages.Landing()))
	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", mux)
}

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"

	mimeTypes := map[string]string{
		".css":   "text/css; charset=utf-8",
		".js":    "application/javascript; charset=utf-8",
		".svg":   "image/svg+xml",
		".html":  "text/html; charset=utf-8",
		".jpg":   "image/jpeg",
		".jpeg":  "image/jpeg",
		".png":   "image/png",
		".gif":   "image/gif",
		".woff":  "font/woff",
		".woff2": "font/woff2",
		".ttf":   "font/ttf",
		".ico":   "image/x-icon",
	}

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)

		if mimeType, ok := mimeTypes[ext]; ok {
			w.Header().Set("Content-Type", mimeType)
		}

		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
