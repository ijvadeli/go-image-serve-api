package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		buf, err := os.ReadFile("img/gopher.png")
		if err != nil {
			http.Error(w, "Image not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(buf)))

		w.Write(buf)
	})

	log.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
