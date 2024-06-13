// package main

// import "fmt"

// func CostCalculator(mealCost float64, taxrate float64, tiprate float64, numPeople int) float64 {

// 	taxAmount := mealCost * (taxrate / 100)
// 	tipAmount := mealCost * (tiprate / 100)
// 	totalCost := mealCost + taxAmount + tipAmount
// 	CostPerPerson := totalCost / float64(numPeople)
// 	return CostPerPerson
// }

// func CostExe() {

// 	var mealCost, tiprate, taxrate float64
// 	var numPeople int

// 	fmt.Println("Please Enter Meal Cost: ")
// 	fmt.Scanf("%f", &mealCost)

// 	fmt.Println("Please Enter Tip: ")
// 	fmt.Scanf("%f", &tiprate)

// 	fmt.Println("Please Enter Tax Rate: ")
// 	fmt.Scanf("%f", &taxrate)

// 	fmt.Println("Please Enter Number Of People: ")
// 	fmt.Scanf("%d", &numPeople)

// 	CostPerPerson := CostCalculator(mealCost, tiprate, taxrate, numPeople)
// 	fmt.Println(CostPerPerson)

// }

package main

import "fmt"

//CostCalculator exported to main
func CostCalculator(mealCost float64, taxRate float64, tipRate float64, numPeople int) float64 {
	taxAmount := mealCost * (taxRate / 100)

	tipAmount := mealCost * (tipRate / 100)

	totalCost := mealCost + taxAmount + tipAmount

	costPerPerson := totalCost / float64(numPeople)

	return costPerPerson
}

//CostExe exported to main
func CostExe() {

	var mealCost, tipRate, taxRate float64
	var numPeople int

	fmt.Println("Please Enter Meal Cost: ")
	fmt.Scanf("%f", &mealCost)

	fmt.Println("Please Enter Tip Rate (as a percentage): ")
	fmt.Scanf("%f", &tipRate)

	fmt.Println("Please Enter Tax Rate (as a percentage): ")
	fmt.Scanf("%f", &taxRate)

	fmt.Println("Please Enter Number Of People: ")
	fmt.Scanf("%d", &numPeople)

	costPerPerson := CostCalculator(mealCost, taxRate, tipRate, numPeople)
	fmt.Printf("The cost per person is: $%.2f\n", costPerPerson)
}
