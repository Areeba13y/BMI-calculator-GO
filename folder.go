package main

import "fmt"

func main() {

	var weight, height, bmi float64

	fmt.Print("Enter your weight (in kg): ")
	fmt.Scan(&weight)

	fmt.Print("Enter your height (in meters): ")
	fmt.Scan(&height)

	bmi = weight / (height * height)

	fmt.Printf("Your BMI is: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Category: Normal weight")
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}
