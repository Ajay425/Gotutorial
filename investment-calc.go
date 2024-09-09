package main  // tells go this is where the execution starts from

import (
	"fmt"
	"math"
)

func main()  {
	const inflationRate = 2.5
	var inveestmentAmount float64 = 2000
	var expectedReturn = 5.5
	var years float64 = 5


	var future = inveestmentAmount * math.Pow(1 + expectedReturn / 100, years)
	inflationValue := future / math.Pow(1 + inflationRate / 100, years)
	
	fmt.Println("Future Value: ",future)
	fmt.Println("Inflation Value: ",inflationValue)
}

/*
Some things to know are:
 - CANNOT USE single quotes in GO while passing a string, Can use `this type of string `
 - Every GO code needs package
 - Package vs Modules
 - Two files cannot have a main func in the same package
*/