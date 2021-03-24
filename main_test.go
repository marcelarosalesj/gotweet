package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect. Got %d and expected %d", total, 10)
	}

}
