package main

import "testing"

func TestHello(t *testing.T) {
    if Hello() != "Hello, world!" {
	t.Errorf("failed")
    }
}
