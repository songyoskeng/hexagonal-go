package begin

type TaxCalculator struct {
	ForCalculateTax
}

func NewTaxCalculator() *TaxCalculator {
	return &TaxCalculator{}
}

func (c *TaxCalculator) CalculateTax(salary float64) float64 {
	taxRepository := &TaxRateRepository{}
	taxRate := taxRepository.GetTaxRateFrom(salary)
	return salary * taxRate
}
