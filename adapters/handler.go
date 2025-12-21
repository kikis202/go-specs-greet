package go_specs_greet

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/kikis202/go-specs-greet/domain-logic"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, go_specs_greet.Greet(name))
}
