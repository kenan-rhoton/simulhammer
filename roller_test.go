package main

import "testing"

func check(t *testing.T, res Results, mindice, maxdice int) {
	count := len(res.Dice)
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
		check(t, Roll(5).Target(4), 0, 5)
	}
}

func TestRollExplosionBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		check(t, Roll(i).Explode(2), i*2, i*2)
	}
}

func TestRollExplosionFunctionBoundaries(t *testing.T) {
	for i := 0; i < 100; i++ {
		fn := func (i int) int {return i}
		check(t, Roll(i).ExplodeFn(fn), i, i*6)
	}
}

func TestRerollCount(t *testing.T) {
	for i := 0; i < 100; i++ {
		check(t, Roll(i).Reroll(), i, i)
	}
}

