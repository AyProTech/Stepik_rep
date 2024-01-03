package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func printError(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}
func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		return
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Println("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}
