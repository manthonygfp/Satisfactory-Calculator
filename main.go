package main

import (
	"fmt"
)

// TODO:
// * Tally up by-products
// * Recipe selection
// * Figure out how to show mining requirements by ore/min instead of buildings (with static size and purity)

type Totals map[string]float64

type Recipe struct {
	Name   string
	Time   float64 // how long it takes for a cycle to complete
	Amount float64 // how many are produced each cycle
	Needs  []Need  // what Products are used to build this Product
}

type Need struct {
	Recipe Recipe
	Amount float64 // how many products are required to fulfill this need
}

var (
	AdaptiveControlUnit = Recipe{
		"Adaptive Control Unit",
		120,
		2,
		[]Need{
			{AutomatedWiring, 15},
			{CircuitBoard, 10},
			{HeavyModularFrame, 2},
			{Computer, 2},
		},
	}

	AutomatedWiring = Recipe{
		"Automated Wiring",
		24,
		1,
		[]Need{
			{Stator, 1},
			{Cable, 20},
		},
	}

	CircuitBoard = Recipe{
		"Circuit Board",
		8,
		1,
		[]Need{
			{CopperSheet, 2},
			{Plastic, 4},
		},
	}

	HeavyModularFrame = Recipe{
		"Heavy Modular Frame",
		30,
		1,
		[]Need{
			{ModularFrame, 5},
			{SteelPipe, 15},
			{EncasedIndustrialBeam, 5},
			{Screw, 100},
		},
	}

	Computer = Recipe{
		"Computer",
		24,
		1,
		[]Need{
			{CircuitBoard, 10},
			{Cable, 9},
			{Plastic, 18},
			{Screw, 52},
		},
	}

	Stator = Recipe{
		"Stator",
		12,
		1,
		[]Need{
			{SteelPipe, 3},
			{Wire, 8},
		},
	}

	Cable = Recipe{
		"Cable",
		2,
		1,
		[]Need{
			{Wire, 2},
		},
	}

	CopperSheet = Recipe{
		"Copper Sheet",
		6,
		1,
		[]Need{
			{CopperIngot, 2},
		},
	}

	Plastic = Recipe{ // By-product not configured
		"Plastic",
		6,
		2,
		[]Need{
			{CrudeOil, 3},
		},
	}

	ModularFrame = Recipe{
		"Modular Frame",
		60,
		2,
		[]Need{
			{ReinforcedIronPlate, 3},
			{IronRod, 12},
		},
	}

	SteelPipe = Recipe{
		"Steel Pipe",
		6,
		2,
		[]Need{
			{SteelIngot, 3},
		},
	}

	EncasedIndustrialBeam = Recipe{
		"Encased Industrial Beam",
		10,
		1,
		[]Need{
			{SteelBeam, 4},
			{Concrete, 5},
		},
	}

	Screw = Recipe{
		"Screw",
		6,
		4,
		[]Need{
			{IronRod, 1},
		},
	}

	Wire = Recipe{
		"Wire",
		4,
		2,
		[]Need{
			{CopperIngot, 1},
		},
	}

	CopperIngot = Recipe{
		"Copper Ingot",
		2,
		1,
		[]Need{
			{CopperOre, 1},
		},
	}

	CrudeOil = Recipe{ // Oil Extractor Mk. 1
		"Crude Oil",
		60,
		120,
		[]Need{},
	}

	ReinforcedIronPlate = Recipe{
		"Reinforced Iron Plate",
		12,
		1,
		[]Need{
			{IronPlate, 6},
			{Screw, 12},
		},
	}

	IronRod = Recipe{
		"Iron Rod",
		4,
		1,
		[]Need{
			{IronIngot, 1},
		},
	}

	SteelIngot = Recipe{
		"Steel Ingot",
		4,
		3,
		[]Need{
			{IronOre, 3},
			{Coal, 3},
		},
	}

	SteelBeam = Recipe{
		"Steel Beam",
		4,
		1,
		[]Need{
			{SteelIngot, 4},
		},
	}

	Concrete = Recipe{
		"Concrete",
		4,
		1,
		[]Need{
			{Limestone, 3},
		},
	}

	CopperOre = Recipe{ // Miner Mk. 2 Normal
		"Copper Ore",
		60,
		120,
		[]Need{},
	}

	IronPlate = Recipe{
		"Iron Plate",
		6,
		2,
		[]Need{
			{IronIngot, 3},
		},
	}

	IronIngot = Recipe{
		"Iron Ingot",
		2,
		1,
		[]Need{
			{IronOre, 1},
		},
	}

	IronOre = Recipe{ // Miner Mk. 2 Normal
		"Iron Ore",
		60,
		120,
		[]Need{},
	}

	Coal = Recipe{ // Miner Mk. 2 Normal
		"Coal",
		60,
		120,
		[]Need{},
	}

	Limestone = Recipe{ // Miner Mk. 2 Normal
		"Limestone",
		60,
		120,
		[]Need{},
	}
)

func CountNeeds(totals *Totals, recipe Recipe) {

	for _, need := range recipe.Needs {
		period := recipe.Time / need.Recipe.Time               // How many cycles of the needed Product occur during the needing Product's cycle
		periodProductionAmount := period * need.Recipe.Amount  // How many needed Products are produced during the needing Product's cycle
		buildingAmount := need.Amount / periodProductionAmount // How many buildings are required to fulfill the need
		(*totals)[need.Recipe.Name] = (*totals)[need.Recipe.Name] + buildingAmount

		// Recursively tally remaining needs
		if len(need.Recipe.Needs) > 0 {
			CountNeeds(totals, need.Recipe)
		}
	}
}

func main() {
	totals := &Totals{}
	CountNeeds(totals, AdaptiveControlUnit)
	for product, buildingAmount := range *totals {
		fmt.Println(product, buildingAmount)
	}
}
