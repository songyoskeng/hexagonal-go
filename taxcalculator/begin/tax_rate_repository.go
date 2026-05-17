package begin

type TaxRateRepository struct {
	ForGetTaxRate
}

func NewTaxRateRepository() *TaxRateRepository {
	return &TaxRateRepository{}
}

func (r *TaxRateRepository) GetTaxRateFrom(salary float64) float64 {
	if salary < 10001 {
		return 0.1
	} else if salary < 20001 {
		return 0.2
	}

	return 0.3
}
