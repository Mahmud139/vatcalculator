package vatcalculator

import (
	"math"
)
func ExclusiveTax(amount float64, tax float64) float64{
	calculation:=amount*(tax/100)
	return math.Round(calculation)
}

func InclusiveTax(amount float64, tax float64) float64{
	calculation:=amount+amount*(tax/100)
	totalPriceAfterTax:=amount*(amount/calculation)
	totalTax:=amount-totalPriceAfterTax
	return math.Round(totalTax)
}

