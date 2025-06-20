package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("Add(1, 2) should be 3")
	}
}
