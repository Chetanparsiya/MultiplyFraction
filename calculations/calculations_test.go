package calculations

import (
	"testing"
)

type FractionTest struct {
	Numerator, Denominator int
}
type fractionTest struct {
	arg1, arg2, expected Fraction
}

type fractionTestReduce struct {
	arg1, expected Fraction
}

type fractionTestReduceToSimpleFraction struct {
	arg1, arg2, expected1, expected2 Fraction
}

var fractionTests = []fractionTest{
	fractionTest{Fraction{1,8},Fraction{3,5},Fraction{3,40}},
	fractionTest{Fraction{1,4},Fraction{4,2},Fraction{4,8}},
	fractionTest{Fraction{4,5},Fraction{7,8},Fraction{28,40}},
	fractionTest{Fraction{1,0},Fraction{1,8},Fraction{}},
	fractionTest{Fraction{1,0},Fraction{1,0},Fraction{}},
	fractionTest{Fraction{0,0},Fraction{1,2},Fraction{0,0}},
	fractionTest{Fraction{4,5},Fraction{7,9},Fraction{28,45}},
	fractionTest{Fraction{8,4},Fraction{1,19},Fraction{8,76}},
	fractionTest{Fraction{6,4},Fraction{5,20},Fraction{30,80}},
	fractionTest{Fraction{5,7},Fraction{4,5}, Fraction{20,35}},
}

var fractionTestsReduce = []fractionTestReduce{
	fractionTestReduce{Fraction{5,40},Fraction{1,8}},
	fractionTestReduce{Fraction{4,8},Fraction{1,2}},
}

var fractionTestsReduceToSimpleFractions = []fractionTestReduceToSimpleFraction{
	fractionTestReduceToSimpleFraction{Fraction{1,8},Fraction{3,5},Fraction{1,8}, Fraction{3,5}},
	fractionTestReduceToSimpleFraction{Fraction{2,4},Fraction{4,2},Fraction{1,2}, Fraction{2,1}},
	fractionTestReduceToSimpleFraction{Fraction{4,5},Fraction{7,8},Fraction{4,5},Fraction{7,8}},
}
func TestReduceToSimpleFraction(t *testing.T){

	for _, test := range fractionTestsReduceToSimpleFractions{
		if output1, output2 := ReduceToSimpleFraction(test.arg1,test.arg2); output1.Numerator == test.expected1.Numerator && output1.Denominator==test.expected1.Denominator && output2.Numerator == test.expected2.Numerator && output2.Denominator==test.expected2.Denominator {
			t.Logf(" Fraction 1 Output %d/%d equal to expected %d/%d ", output1.Numerator, output1.Denominator,test.expected1.Numerator,test.expected1.Denominator)
			t.Logf(" Fraction 2 Output %d/%d equal to expected %d/%d ", output2.Numerator, output2.Denominator,test.expected2.Numerator,test.expected2.Denominator)
			t.Log("*************************")
		}
	}
}

func TestMultiplyCalculations(t *testing.T){

		for _, test := range fractionTests{
			if output,err := MultiplyCalculations(test.arg1,test.arg2); output.Numerator == test.expected.Numerator && output.Denominator==test.expected.Denominator && err == nil {
				t.Logf("Output %d/%d equal to expected %d/%d ", output.Numerator, output.Denominator,test.expected.Numerator,test.expected.Denominator)
			} else {
				if(err!=nil){
					t.Log(err)
				}else {
					t.Errorf("Output %d/%d not equal to expected %d/%d ", output.Numerator, output.Denominator,test.expected.Numerator,test.expected.Denominator)
				}
			}
	}
}

func TestReduceFunction(t *testing.T){

	for _, test := range fractionTestsReduce{
		if output := ReduceFraction(test.arg1); output.Numerator == test.expected.Numerator && output.Denominator==test.expected.Denominator {
			t.Logf("Output %d/%d equal to expected %d/%d ", output.Numerator, output.Denominator,test.expected.Numerator,test.expected.Denominator)
		} else {
			t.Errorf("Output %d/%d not equal to expected %d/%d ", output.Numerator, output.Denominator,test.expected.Numerator,test.expected.Denominator)
		}
	}
}


