package main

import (
	"errors"
	"fmt"
)

func calculateBMI(weight float64, height float64) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	bmi := weight / (height * height)

	return bmi, nil
}

func main() {
	weight := 70.0
	height := 1.75

	bmi, err := calculateBMI(weight, height)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("weight: %.1f \n", weight)
	fmt.Printf("height: %.2f \n", height)
	fmt.Printf("BMI: %.2f \n", bmi)
}
