package main

import "testing"

func TestHash(t *testing.T) {
	hashed := Hash("angryMonkey")
	const expected = "" +
		"6441e1581eb9814973755c2d0d002b13" +
		"2c7e2952f3a7f69369168f941cd84481" +
		"63eaf8c576a11bd10e41f3354a099d2f" +
		"29b64f664949cf415deecbb603e81fed"
	if hashed != expected {
		t.Errorf("Expected:\n%v\nbut got:\n%v\n", expected, hashed)
	}
}
