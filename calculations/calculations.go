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

func GCD(a, b int,c1 chan int) {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	c1 <- a 
}
func ReduceFraction(f1 Fraction,mychnl chan Fraction) {
	c1 := make(chan int)
	go GCD(f1.Numerator, f1.Denominator,c1)
	gcdValue := <- c1
	f1.Numerator /= gcdValue
	f1.Denominator /= gcdValue
	mychnl <- f1
	close(mychnl)
}

func ReduceToSimpleFraction(f1 Fraction, f2 Fraction) (Fraction, Fraction) {	
	c1 := make(chan Fraction) 
	c2 := make(chan Fraction) 
	go ReduceFraction(f1,c1)
	reducedFraction1 := <- c1
	go ReduceFraction(f2,c2)
	reducedFraction2 := <- c2
	return reducedFraction1, reducedFraction2
}