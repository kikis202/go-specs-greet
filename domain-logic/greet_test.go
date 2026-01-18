package interactions_test

import (
	"testing"

	interactions "github.com/kikis202/go-specs-greet/domain-logic"
	"github.com/kikis202/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
