package calculations

import (
	"errors"
)

type Fraction struct {
	Numerator, Denominator int
}

func MultiplyCalculations(f1 Fraction, f2 Fraction) (Fraction, error) {
	if f1.Numerator ==0 && f1.Denominator ==0 || f2.Numerator==0 && f2.Denominator==0 {
		return Fraction{0,0}, nil
	}
	if f1.Denominator == 0 || f2.Denominator == 0 {
		return Fraction{}, errors.New("denominator Can't contains 0, it cause to infinity")
	}
	f1.Numerator *= f2.Numerator
	f1.Denominator *= f2.Denominator
	if(f1.Numerator==0 ){
		f1.Denominator=0;
	}
	return f1, nil
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func ReduceFraction(f1 Fraction) Fraction {
	gcdValue := GCD(f1.Numerator, f1.Denominator)
	f1.Numerator /= gcdValue
	f1.Denominator /= gcdValue
	return f1
}

func ReduceToSimpleFraction(f1 Fraction, f2 Fraction) (Fraction, Fraction) {	
	reducedFraction1 := ReduceFraction(f1)
	reducedFraction2 := ReduceFraction(f2)
	return reducedFraction1, reducedFraction2
}