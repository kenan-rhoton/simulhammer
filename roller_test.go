package main

import "testing"

func check(t *testing.T, res Results, mindice, maxdice int) {
	count := res.Count()
	if count < mindice {
		t.Errorf("%d is less dice than minimum of %d", count, mindice)
	} else if count > maxdice {
		t.Errorf("%d is more dice than maximum of %d", count, maxdice)
	}

	for _, die := range res.Dice {
		if die < 0 {
			t.Errorf("Wild die is less than 0: %d", die)
		} else if die > 6 {
			t.Errorf("Wild die is more than 6: %d", die)
		}
	}
}

func TestRollCount(t *testing.T) {
	for i := 0; i < 100; i++ {
		check(t, Roll(i), i, i)
	}
}

func TestRollTargetBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		res := Roll(5).Target(4)
		check(t, res, 0, 5)
		for _, die := range res.Dice {
			if die < 4 {
				t.Errorf("Wild die is less than 4: %d", die)
			}
		}
	}
}

func TestRollUnderBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		res := Roll(5).Under(4)
		check(t, res, 0, 5)
		for _, die := range res.Dice {
			if die >= 4 {
				t.Errorf("Wild die is less than 4: %d", die)
			}
		}
	}
}

func TestRollExplosionBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		check(t, Roll(i).Explode(2), i*2, i*2)
	}
}

func TestRollExplosionFunctionBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		fn := func(i int) int { return i }
		check(t, Roll(i).ExplodeFn(fn), i, i*6)
	}
}

func TestRollSumFunctionBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		sum := Roll(i).Sum()
		if sum < i {
			t.Errorf("%d is less dice than minimum of %d", sum, i)
		} else if sum > i*6 {
			t.Errorf("%d is more dice than maximum of %d", sum, i*6)
		}
	}
}

func TestRerollCount(t *testing.T) {
	for i := 0; i < 100; i++ {
		check(t, Roll(i).Reroll(), i, i)
	}
}
