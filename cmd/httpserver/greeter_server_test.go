package main_test

import (
	"testing"

	go_specs_greet "github.com/kikis202/go-specs-greet/adapters/httpserver"
	"github.com/kikis202/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
