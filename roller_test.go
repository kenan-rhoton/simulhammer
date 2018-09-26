package main

import "testing"

func TestRollWithTargetBoundaries(t *testing.T) {
	for i := 0; i < 1000; i++ {
		res := Roll(5).Pass(4).Count()
		if res < 0 || res > 5 {
			t.Errorf("%d is out of [0,5] bounds", res)
		} else {
			t.Logf("Roll 5 dice with target 4 result: %d", res)
		}
	}
}
