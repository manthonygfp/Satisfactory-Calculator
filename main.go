package main

import (
	"fmt"
)

// TODO:
// * Tally up by-products

type RecipeBook struct {
	Default    Recipe
	Alternates map[string]Recipe // The map keys are the Recipe names
}

type Product struct {
	Recipes RecipeBook
}

type Recipe struct {
	Time        float64      // How long it takes for a cycle to complete
	Amount      float64      // How many Products are produced each cycle
	Ingredients []Ingredient // What Ingredients are used in this Recipe
}

type Ingredient struct {
	ProductName string
	Amount      float64 // How many products are required to fulfill this need
}

type Tally struct {
	Extracted map[string]float64
	Small     map[string]float64
	Medium    map[string]float64
	Large     map[string]float64
}

var (
	ActiveAlternates = map[string]string{
		"Screw":    "Cast Screw",
		"Iron Rod": "Steel Rod",
	}

	Products = map[string]Product{
		// Mined
		"Coal":       {RecipeBook{Recipe{0, 0, []Ingredient{}}, map[string]Recipe{}}},
		"Copper Ore": {RecipeBook{Recipe{0, 0, []Ingredient{}}, map[string]Recipe{}}},
		"Iron Ore":   {RecipeBook{Recipe{0, 0, []Ingredient{}}, map[string]Recipe{}}},
		"Limestone":  {RecipeBook{Recipe{0, 0, []Ingredient{}}, map[string]Recipe{}}},

		// Extracted
		"Crude Oil": {RecipeBook{Recipe{0, 0, []Ingredient{}}, map[string]Recipe{}}},

		// Smelted
		"Concrete": {RecipeBook{
			Recipe{4, 1, []Ingredient{
				{"Limestone", 3},
			}},
			map[string]Recipe{},
		}},

		"Copper Ingot": {RecipeBook{
			Recipe{2, 1, []Ingredient{
				{"Copper Ore", 1},
			}},
			map[string]Recipe{},
		}},

		"Iron Ingot": {RecipeBook{
			Recipe{2, 1, []Ingredient{
				{"Iron Ore", 1},
			}}, map[string]Recipe{},
		}},

		// Foundried (lol)
		"Steel Ingot": {RecipeBook{
			Recipe{4, 3, []Ingredient{
				{"Iron Ore", 3},
				{"Coal", 3},
			}},
			map[string]Recipe{},
		}},

		// Constructed
		"Cable": {RecipeBook{
			Recipe{2, 1, []Ingredient{
				{"Wire", 2},
			}},
			map[string]Recipe{},
		}},

		"Copper Sheet": {RecipeBook{
			Recipe{6, 1, []Ingredient{
				{"Copper Ingot", 2},
			}},
			map[string]Recipe{},
		}},

		"Iron Plate": {RecipeBook{
			Recipe{6, 2, []Ingredient{
				{"Iron Ingot", 3},
			}}, map[string]Recipe{},
		}},

		"Iron Rod": {RecipeBook{
			Recipe{4, 1, []Ingredient{
				{"Iron Ingot", 1},
			}},
			map[string]Recipe{
				"Steel Rod": {5, 4, []Ingredient{
					{"Steel Ingot", 1},
				}},
			},
		}},

		"Screw": {RecipeBook{
			Recipe{6, 4, []Ingredient{
				{"Iron Rod", 1},
			}},
			map[string]Recipe{
				"Cast Screw": {24, 20, []Ingredient{
					{"Iron Ingot", 5},
				}},
			},
		}},

		"Steel Beam": {RecipeBook{
			Recipe{4, 1, []Ingredient{
				{"Steel Ingot", 4},
			}},
			map[string]Recipe{},
		}},

		"Steel Pipe": {RecipeBook{
			Recipe{6, 2, []Ingredient{
				{"Steel Ingot", 3},
			}},
			map[string]Recipe{},
		}},

		"Wire": {RecipeBook{
			Recipe{4, 2, []Ingredient{
				{"Copper Ingot", 1},
			}},
			map[string]Recipe{},
		}},

		// Assembled
		"Automated Wiring": {RecipeBook{
			Recipe{24, 1, []Ingredient{
				{"Stator", 1},
				{"Cable", 20},
			}},
			map[string]Recipe{},
		}},

		"Circuit Board": {RecipeBook{
			Recipe{8, 1, []Ingredient{
				{"Copper Sheet", 2},
				{"Plastic", 4},
			}},
			map[string]Recipe{},
		}},

		"Encased Industrial Beam": {RecipeBook{
			Recipe{10, 1, []Ingredient{
				{"Steel Beam", 4},
				{"Concrete", 5},
			}},
			map[string]Recipe{},
		}},

		"Modular Frame": {RecipeBook{
			Recipe{60, 2, []Ingredient{
				{"Reinforced Iron Plate", 3},
				{"Iron Rod", 12},
			}},
			map[string]Recipe{},
		}},

		"Reinforced Iron Plate": {RecipeBook{
			Recipe{12, 1, []Ingredient{
				{"Iron Plate", 6},
				{"Screw", 12},
			}},
			map[string]Recipe{},
		}},

		"Stator": {RecipeBook{
			Recipe{12, 1, []Ingredient{
				{"Steel Pipe", 3},
				{"Wire", 8},
			}},
			map[string]Recipe{},
		}},

		// Manufactured
		"Adaptive Control Unit": {RecipeBook{
			Recipe{120, 2, []Ingredient{
				{"Automated Wiring", 15},
				{"Circuit Board", 10},
				{"Heavy Modular Frame", 2},
				{"Computer", 2},
			}},
			map[string]Recipe{},
		}},

		"Computer": {RecipeBook{
			Recipe{24, 1, []Ingredient{
				{"Circuit Board", 10},
				{"Cable", 9},
				{"Plastic", 18},
				{"Screw", 52},
			}},
			map[string]Recipe{},
		}},

		"Heavy Modular Frame": {RecipeBook{
			Recipe{8, 1, []Ingredient{
				{"Modular Frame", 5},
				{"Steel Pipe", 15},
				{"Encased Industrial Beam", 5},
				{"Screw", 100},
			}},
			map[string]Recipe{},
		}},

		// Refined
		"Plastic": {RecipeBook{
			Recipe{6, 1, []Ingredient{
				{"Copper Ingot", 2},
			}},
			map[string]Recipe{},
		}},
	}
)

func PreferRecipe(productName string) Recipe {
	recipe := Products[productName].Recipes.Default
	if val, ok := ActiveAlternates[productName]; ok {
		recipe = Products[productName].Recipes.Alternates[val]
	}
	return recipe
}

func TallyIngredients(productName string, tally *Tally) {
	recipe := PreferRecipe(productName)

	for _, ingredient := range recipe.Ingredients {
		ingredientRecipe := PreferRecipe(ingredient.ProductName)

		if len(ingredientRecipe.Ingredients) == 0 {
			(*tally).Extracted[ingredient.ProductName] = (*tally).Extracted[ingredient.ProductName] + (60 / recipe.Time * ingredient.Amount)
		} else {
			period := recipe.Time / ingredientRecipe.Time                // How many cycles of the needed Recipe occur during the needing Recipe's cycle
			periodProductionAmount := period * ingredientRecipe.Amount   // How many Ingredients are produced during the needing Recipe's cycle
			buildingAmount := ingredient.Amount / periodProductionAmount // How many buildings are required to fulfill the need
			switch len(ingredientRecipe.Ingredients) {
			case 1:
				(*tally).Small[ingredient.ProductName] = (*tally).Small[ingredient.ProductName] + buildingAmount
			case 2:
				(*tally).Medium[ingredient.ProductName] = (*tally).Medium[ingredient.ProductName] + buildingAmount
			case 3, 4:
				(*tally).Large[ingredient.ProductName] = (*tally).Large[ingredient.ProductName] + buildingAmount
			}
		}

		TallyIngredients(ingredient.ProductName, tally)
	}
}

func main() {
	tally := &Tally{map[string]float64{}, map[string]float64{}, map[string]float64{}, map[string]float64{}} // TODO: ew
	TallyIngredients("Adaptive Control Unit", tally)

	fmt.Println("Amount needed to be extracted per minute:")
	for product, amount := range *&tally.Extracted {
		fmt.Println(product, amount)
	}
	fmt.Println()

	fmt.Println("Building amount required for processing:")
	sections := []map[string]map[string]float64{
		{"One ingredient: ": tally.Small},
		{"Two ingredients: ": tally.Medium},
		{"Three or more ingredients: ": tally.Large},
	}
	for _, section := range sections {
		for title, tally := range section {
			fmt.Println(title)

			for product, amount := range tally {
				fmt.Println(product, amount)
			}

			fmt.Println()
		}
	}
}
