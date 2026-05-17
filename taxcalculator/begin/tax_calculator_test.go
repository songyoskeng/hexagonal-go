package begin

import "testing"

func TestCalculateTax(t *testing.T) {
	taxCalculator := &TaxCalculator{}
	tax := taxCalculator.CalculateTax(10000)
	if tax != 1000 {
		t.Fatalf("expected 1000, got %v", tax)
	}
}
