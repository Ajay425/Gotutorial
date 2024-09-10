package main // tells go this is where the execution starts from

import (
	"fmt"
	"math"

)

const inflationRate = 2.5

func main()  {
	var investmentAmount float64 = 2000
	var expectedReturn = 5.5
	var years float64 = 5


	 // var future = investmentAmount * math.Pow(1 + expectedReturn / 100, years)
	 // inflationValue := future / math.Pow(1 + inflationRate / 100, years)
	

	future, inflationValue := calculateFuctureValues(investmentAmount, expectedReturn, years)
	
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", future)
	formattedRFV := fmt.Sprintf("Inflation Value (Adjusted for Inflation): %.1f\n", inflationValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string){
	// // put the code on what the function does
	fmt.Print(text)
}

func calculateFuctureValues(investmentAmount float64, expectedReturn float64, years float64) (fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1 + expectedReturn / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return
}


/*
Some things to know are:
 - CANNOT USE single quotes in GO while passing a string, Can use `this type of string `
 - Every GO code needs package
 - Package vs Modules
 - Two files cannot have a main func in the same package
 - Using backticks allows you to break the string in multiple lines and print it out vs Double quotes you cant do that
*/