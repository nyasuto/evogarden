package world

import (
	"math/rand"
	"testing"
)

func TestEvolveImprovesEnergy(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	baseline := Genome{Vision: 1, MoveCost: 3, FoodGain: 1}
	baseEnergy := evaluateGenome(r, baseline)

	best := Evolve(rand.New(rand.NewSource(42)), 50, 20)
	evalEnergy := evaluateGenome(rand.New(rand.NewSource(99)), best)
	if evalEnergy <= baseEnergy {
		t.Fatalf("expected evolved genome energy %d to be greater than baseline %d", evalEnergy, baseEnergy)
	}
}
