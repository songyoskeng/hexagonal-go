package solution

type TaxCalculator struct {
	ForCalculateTax
	forGetTaxRate ForGetTaxRate
}

func NewTaxCalculator(forGetTaxRate ForGetTaxRate) *TaxCalculator {
	return &TaxCalculator{forGetTaxRate: forGetTaxRate}
}

func (c *TaxCalculator) CalculateTax(salary float64) float64 {
	taxRate := c.forGetTaxRate.GetTaxRateFrom(salary)
	return salary * taxRate
}
