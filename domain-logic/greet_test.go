package go_specs_greet_test

import (
	"testing"

	go_specs_greet "github.com/kikis202/go-specs-greet/domain-logic"
	"github.com/kikis202/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))
}
