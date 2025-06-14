package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationrate float64 = 2.5
	var investmetAmount float64 //you can leave a var without assigning a value but go automatically assign it to 0 (null)
	expectedReturnRate := 5.5
	var years float64 = 10 //we should have a type casted if we do not assign a value to a variable.

	fmt.Print("Enter your investment amount : ")
	fmt.Scan(&investmetAmount)

	fmt.Print("Expected return rate : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of investing years : ")
	fmt.Scan(&years)

	futureValue := investmetAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationrate/100, years)
	fmt.Println(futureValue)
	fmt.Print(futureRealValue)
	//for type conversion float64() can be used
}
