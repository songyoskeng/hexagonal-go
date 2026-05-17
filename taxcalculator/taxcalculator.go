package taxcalculator

import "hexagonalgo/taxcalculator/solution"

type TaxCalculator = solution.TaxCalculator
type TaxRepository = solution.TaxRateRepository
type ForGetTaxRate = solution.ForGetTaxRate

func NewTaxCalculator(forGetTaxRate solution.ForGetTaxRate) *TaxCalculator {
	return solution.NewTaxCalculator(forGetTaxRate)
}

func NewTaxRateRepository() *TaxRepository {
	return solution.NewTaxRateRepository()
}
