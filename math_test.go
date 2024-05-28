package main

import "testing"

func TestSoma(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Errorf("Error")
	}

}
