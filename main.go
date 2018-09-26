package main

import (
	"fmt"
	"math/rand"
	"time"
)

func LightningHammers(qty int) int {
	res := 0
	hit := Roll(qty).Target(3)
	res += hit.Target(6).Count() * 2
	wounds := hit.Under(6).Reroll().Target(3)
	res += wounds.Reroll().Under(4-1).Count() * 2
	return res
}

func StarsoulMaces(qty int) int {
	res := 0
	hit := Roll(qty).Target(2)
	for i := 0; i < hit.Under(6).Count(); i++ {
		res += D3()
	}
	for i := 0; i < hit.Target(6).Count(); i++ {
		res += D3() + 1
	}
	return res
}

func Grandstaves(qty int) int {
	res := 0
	hit := Roll(qty).Target(3)
	wounds := hit.Reroll().Target(3)
	res += wounds.Reroll().Under(4).Count() * 2
	return res
}

func CelestialLightningArcs(qty int) int {
	return Roll(qty).Target(4).Count()
}

func Ret_v_Evo() int {
	return LightningHammers(7) + StarsoulMaces(2)
}

func Evo_v_Ret() int {
	return Grandstaves(16) + CelestialLightningArcs(10)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Retributors vs Evocators: %d\n", Ret_v_Evo())
	fmt.Printf("Evocators vs Retributors: %d\n", Evo_v_Ret())
}
