package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var (
	//go:embed react/my-app/build
	web embed.FS
)

const defaultPort = "8080"

func main() {
	dist, _ := fs.Sub(web, "react/my-app/build")

	http.Handle(
		"/", http.FileServer(http.FS(dist)),
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	fmt.Printf("O aplicativo está disponível em http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
