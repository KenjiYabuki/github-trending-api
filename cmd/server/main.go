package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KenjiYabuki/github-trending-api/internal/app"
)

func main() {
	http.HandleFunc("/repos", app.ReposHandler)
	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
