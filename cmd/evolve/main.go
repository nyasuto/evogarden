package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"evogarden/world"
)

func main() {
	gens := flag.Int("gens", 1000, "number of generations")
	pop := flag.Int("pop", 100, "population size")
	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	best := world.Evolve(r, *gens, *pop)

	// run a demonstration simulation with agents using the best genome
	g := world.NewGrid(20, 20)
	_ = g.Set(10, 10, world.Food)
	agents := []*world.Agent{
		{X: 0, Y: 0, Grid: g, Energy: 20, Vision: best.Vision, MoveCost: best.MoveCost, FoodGain: best.FoodGain},
		{X: 19, Y: 19, Grid: g, Energy: 20, Vision: best.Vision, MoveCost: best.MoveCost, FoodGain: best.FoodGain},
	}
	sim := &world.Simulation{Grid: g, Agents: agents}
	for step := 0; step < 200; step++ {
		sim.Step()
		// fmt.Printf("Step %d:\n%s\n\n", step+1, world.Render(g, sim.Agents))
	}

	fmt.Printf("Best genome after %d generations: Vision=%d MoveCost=%d FoodGain=%d\n",
		*gens, best.Vision, best.MoveCost, best.FoodGain)

}
