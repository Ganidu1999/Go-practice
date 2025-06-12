package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationrate float64 = 2.5
	var investmetAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	futureValue := investmetAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationrate/100, years)
	fmt.Println(futureValue)
	fmt.Print(futureRealValue)
	//for type conversion float64() can be used
}
