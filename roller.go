package main

import "math/rand"

func RollWithTarget(dice, target int) int {
	res := 0
	for i := 0; i < dice; i++ {
		if rand.Intn(7) >= target {
			res += 1
		}
	}
	return res
}
