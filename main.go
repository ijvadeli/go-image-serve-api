package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Distribution", `attachment; filename="img/gopher.png"`)
		http.ServeFile(w, r, "img/gopher.png")
	})

	log.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
