package main

import (
	"fmt"
	"log"
	"net/http"
)

// Composable web libraries / framewoprks:
// - https://github.com/go-chi/chi
// - https://github.com/go-ozzo/ozzo-routing
// Frameworks:
// - https://github.com/gin-gonic/gin
// - https://github.com/kataras/iris

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
