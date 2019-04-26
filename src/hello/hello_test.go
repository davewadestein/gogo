package main

import (
	"testing"
	"gogo"
)

func TestHello(t *testing.T) {
    if gogo.Hello() != "Hello, world!" {
	t.Errorf("failed")
    }
}
