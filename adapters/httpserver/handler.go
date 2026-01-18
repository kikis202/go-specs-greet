package httpserver

import (
	"fmt"
	"net/http"

	interactions "github.com/kikis202/go-specs-greet/domain-logic"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, interactions.Greet(name))
}
