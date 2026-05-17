package solution

import "testing"

func TestTaxCalculator(t *testing.T) {
	fromTaxRateRepository := &TaxRateRepository{}
	taxCalculator := NewTaxCalculator(fromTaxRateRepository)
	tax := taxCalculator.CalculateTax(10000)

	if tax != 1000 {
		t.Fatalf("expected 1000, got %v", tax)
	}
}
