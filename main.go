package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(4),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.Value)
	fmt.Printf("Nutrition Grade: %s\n", ns.GetNutritionGrade())
}
