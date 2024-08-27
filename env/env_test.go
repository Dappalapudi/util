package env

import (
	"testing"
)

func TestGet(t *testing.T) {
	if Test != Get() {
		t.Errorf("expecting different default environment: got %v want %v", Get(), Development)
	}
}

func TestIs(t *testing.T) {
	table := []struct {
		envir  Environment
		expect bool
	}{
		{Development, true},
		{Stage, false},
		{Production, false},
	}

	Set(Development)

	for _, a := range table {
		if Is(a.envir) != a.expect {
			t.Errorf("for %q, got %v want %v", a.envir, !a.expect, a.expect)
		}
	}
}
