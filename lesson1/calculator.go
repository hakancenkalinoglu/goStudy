package main

import "fmt"

// Calculate fonksiyonu artık dışarıdan 3 tane değer alıyor ve geriye float64 tipinde bir sonuç döndürüyor.
func Calculate(num1 float64, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		// Bölme işleminde sıfıra bölünme hatasını önlemek için küçük bir kontrol ekledik
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Cannot divide by zero!")
			result = 0
		}
	default:
		fmt.Println("Error: Invalid operator.")
		result = 0
	}

	// Hesaplanan sonucu fonksiyondan dışarıya "fırlatıyoruz"
	return result
}
