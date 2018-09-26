package main

import "math/rand"

type Results struct {
	Dice []int
}

func D3() int {
	return rand.Intn(3) + 1
}

func D6() int {
	return rand.Intn(6) + 1
}

func Roll(dice int) Results {
	res := make([]int, dice)
	for die := range res {
		res[die] = D6()
	}
	return Results{Dice: res}
}

func (r Results) Target(target int) Results {
	res := make([]int, 0)
	for _, die := range r.Dice {
		if die >= target {
			res = append(res, die)
		}
	}
	return Results{Dice: res}
}

func (r Results) Under(target int) Results {
	res := make([]int, 0)
	for _, die := range r.Dice {
		if die < target {
			res = append(res, die)
		}
	}
	return Results{Dice: res}
}

func (r Results) ExplodeFn(fn func(int) int) Results {
	res := make([]int, 0)
	for _, die := range r.Dice {
		times := fn(die)
		for i := 0; i < times; i++ {
			res = append(res, die)
		}
	}
	return Results{Dice: res}
}

func (r Results) Explode(n int) Results {
	return r.ExplodeFn(func(_ int) int { return n })
}

func (r Results) Reroll() Results {
	res := make([]int, len(r.Dice))

	for i := range r.Dice {
		res[i] = D6()
	}
	return Results{Dice: res}
}

func (r Results) Count() int {
	return len(r.Dice)
}

func (r Results) Sum() int {
	res := 0
	for i := range r.Dice {
		res += r.Dice[i]
	}
	return res
}
