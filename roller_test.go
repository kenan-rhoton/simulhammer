package main

import "testing"

func TestRollBoundariesStatistics(t *testing.T) {
	for i := 0; i < 1000; i++ {
		res := RollWithTarget(5, 4)
		if res < 0 || res > 5 {
			t.Errorf("%d is out of [0,5] bounds", res)
		} else {
			t.Logf("Result: %d", res)
		}
	}
}
