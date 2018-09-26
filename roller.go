package main

import "math/rand"

type Results struct{
	dice []int
}

func d6() int {
	return rand.Intn(7)
}

func Roll(dice int) Results {
	res := make([]int, dice)
	for die := range res {
		res[die] = d6()
	}
	return Results{dice: res}
}

func (r Results) Pass(target int) Results {
	res := make([]int, 0)
	for _, die := range r.dice {
		if die >= target {
			res = append(res, die)
		}
	}
	return Results{dice: res}
}

func (r Results) Count() int {
	return len(r.dice)
}

