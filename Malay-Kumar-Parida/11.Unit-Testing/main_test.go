package main

import "testing"

func TestCalculate(t *testing.T){ // Single test case
	if Calculate(2)!=4{
		t.Error("Expected result was 4")
	}
}

type TestCase struct{
	input int
	expected int
}

//Positive Testing
func TestTableCalculate(t* testing.T){ // Test table where we can test a bunch of test cases in a slice
	
	var tests=[]TestCase{
		{2,4},
		{-1,1},
		{100000,100002},
		{0,2},
	}


	for _,test :=range tests{
		if output :=Calculate(test.input);output!=test.expected{
			t.Error("Test Failed:{} inputted,{} expected, received:{}",test.input,test.expected,output)
		}
	}
}

//Negative Testing
func TestNegativeTableCalculate(t* testing.T){ // Test table where we can test a bunch of test cases in a slice
	
	var tests=[]TestCase{
		{2,5},
		{-1,0},
		{100000,100004},
		{0,5},
	}


	for _,test :=range tests{
		if output :=Calculate(test.input);output==test.expected{ // Negative testing
			t.Error("Test Failed:{} inputted,{} expected, received:{}",test.input,test.expected,output)
		}
	}
}

// Can do go test --cover on terminal to check statement coverage