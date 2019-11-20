package newmod

import (
	"log"
	"testing"

	"github.com/huangc28/newmod/world"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	log.Printf("world %s", world.World())
}
