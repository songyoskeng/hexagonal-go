package begin

type TaxCalculator struct{}

func (c *TaxCalculator) CalculateTax(salary float64) float64 {
	taxRepository := &TaxRateRepository{}
	taxRate := taxRepository.GetTaxRateFrom(salary)
	return salary * taxRate
}
