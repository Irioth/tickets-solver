package main

import "testing"

func TestPowerReverse(t *testing.T) {
	v, ok := power.ReverseA(100, 2)
	if !ok {
		t.Log("not applicable")
		t.Fail()
	}
	if v != 10 {
		t.Log("wrong value")
		t.Fail()
	}
}
