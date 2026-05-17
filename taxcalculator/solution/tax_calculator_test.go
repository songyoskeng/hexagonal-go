package solution

import "testing"

func TestTaxCalculator(t *testing.T) {
	taxRepository := NewTaxRateRepository()
	taxCalculator := NewTaxCalculator(taxRepository)
	tax := taxCalculator.CalculateTax(10000)

	if tax != 1000 {
		t.Fatalf("expected 1000, got %v", tax)
	}
}
