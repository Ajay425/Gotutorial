package main

import (
	"fmt"

)



func main() {

	revenue := askInput("Revenue: ")
	expenses:= askInput("Expenses: ")
	taxRate := askInput("Tax Rate: ")

	//ebt := revenue - expenses
	//profit := ebt * (1 - taxRate/100)
	//ratio := ebt / profit
	
	ebt, profit, ratio := calculateTax(revenue, expenses, taxRate)


	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func askInput(text string) float64{
	var userInput float64
	fmt.Print(text) 
	fmt.Scan(&userInput)

	return userInput
}

func calculateTax(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
