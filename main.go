package main

import (
	"fmt"

	"fraction-multiply/calculations"
)

func main() {
	var numerator1,testcase, numerator2,denominator1,denominator2 int
	fmt.Println("Enter the number of Test Cases ")
	fmt.Scan(&testcase)
	for i:=0 ;i<testcase;i++ {
		for {
			fmt.Println("Enter First Fraction first Numerator then Denominator")
			fmt.Scan(&numerator1,&denominator1)
			fmt.Println("Enter Second Fraction first Numerator then Denominator")
			fmt.Scan(&numerator2,&denominator2)
			if denominator1!=0 && denominator2!=0 {
					break
			}else {
				fmt.Println("Enter Denominator Greater than 0")
			}
	}
		
		fraction1 := calculations.Fraction{numerator1,denominator1}
		fraction2 := calculations.Fraction{numerator2,denominator2}
		c := make(chan calculations.Fraction) 
	
		var reducedFraction1, reducedFraction2 = calculations.ReduceToSimpleFraction(fraction1,fraction2)
		fmt.Printf("Simple Fraction 1 is %d/%d\n",reducedFraction1.Numerator,reducedFraction1.Denominator)
		fmt.Printf("Simple Fraction 2 is %d/%d\n",reducedFraction2.Numerator,reducedFraction2.Denominator)
		result,err := calculations.MultiplyCalculations(fraction1, fraction2)
		if(result.Numerator!=0){
			if(err!=nil){
				fmt.Print("Invalid Inputs ",err.Error())
			} else {
				fmt.Printf("Result of Multiplications of %d/%d * %d/%d is %d/%d\n",reducedFraction1.Numerator,reducedFraction1.Denominator,reducedFraction2.Numerator,reducedFraction2.Denominator,result.Numerator,result.Denominator)
				go calculations.ReduceFraction(result,c)
				res := <- c
				fmt.Printf("Result of Multiplications in simplest Form %d/%d\n",res.Numerator,res.Denominator)	
			}
		} else {
			//fmt.Printf("Result of Multiplications of %d/%d * %d/%d is %d/%d\n",reducedFraction1.Numerator,reducedFraction1.Denominator,reducedFraction2.Numerator,reducedFraction2.Denominator,result.Numerator,result.Denominator)
		}
	}
	
}