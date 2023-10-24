package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)

	if total != 10 {
		t.Errorf("The number %v does not equal to 10", total)
	}

}
