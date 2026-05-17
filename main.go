package main

import (
	"fmt"
	"hexagonalgo/taxcalculator"
)

func main() {
	fmt.Println("---Tax calculator---")
	taxRepo := taxcalculator.NewTaxRateRepository()
	taxApp := taxcalculator.NewTaxCalculator(taxRepo)

	taxResult := taxApp.CalculateTax(2000)
	fmt.Printf("result: %.2f\n", taxResult)

}
