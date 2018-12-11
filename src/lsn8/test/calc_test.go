package main

import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add error")
	}
	t.Logf("succeed")
}
