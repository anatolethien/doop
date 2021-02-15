package main

import (
	"fmt"
	"os"
)

const Max = 9223372036854775807
const Min = -9223372036854775808

func main() {

	if argsValid(os.Args) == false {
		return
	}

	var a, operator, b = Atoi(os.Args[1]), os.Args[2], Atoi(os.Args[3])
	var result int

	if operator == "/" && b == 0 {
		fmt.Println("No division by 0")
		return
	}

	if operator == "%" && b == 0 {
		fmt.Println("No modulo by 0")
		return
	}

	switch operator {
	case "+":
		result = addition(a, b)
	case "-":
		result = subtraction(a, b)
	case "*":
		result = multiplication(a, b)
	case "/":
		result = division(a, b)
	case "%":
		result = modulo(a, b)
	}

	fmt.Println(result)

}

func argsValid(args []string) bool {

	if len(args) != 4 {
		return false
	}

	if isNumeric(args[1]) == false {
		return false
	}

	if isOperator(args[2]) == false {
		return false
	}

	if isNumeric(args[3]) == false {
		return false
	}

	return true
}

func isNumeric(number string) bool {

	for _, v := range number {
		if v != 45 && (v < 48 || 57 < v) {
			return false
		}
	}

	return true
}

func isOperator(operator string) bool {

	switch operator {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case "%":
		return true
	default:
		return false
	}

}

func Atoi(s string) int {

	var f = false
	var signCount, str = 0, 0

	for i, v := range s {
		// 43 = '+', 45 = '-', 48 = '0', 57 = '9'
		if v == 43 || v == 45 || v >= 48 && v <= 57 {

			if (v == 43 || v == 45) && i > 0 {
				return 0
			}

			if v == 45 && i == 0 {
				f = true
			}

			if v == 43 || v == 45 {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}

			var a = 0
			var j int32
			for j = 49; j <= v; j++ {
				a++
			}
			str = str*10 + a

		} else {
			return 0
		}

	}

	if f == true {
		str = str - (str * 2)
	}

	return str
}

func addition(a, b int) int {
	if ((b > 0) && (a > (Max - b))) || ((b < 0) && (a < (Min - b))) {
		return 0
	}
	return a + b
}

func subtraction(a, b int) int {
	if (b > 0 && a < Min+b) || (b < 0 && a > Max+b) {
		return 0
	}
	return a - b
}

func multiplication(a, b int) int {
	if a > 0 {
		if b > 0 {
			if a > (Max / b) {
				return 0
			}
		} else {
			if b < (Min / a) {
				return 0
			}
		}
	} else {
		if b > 0 {
			if a < (Min / b) {
				return 0
			}
		} else {
			if a != 0 && (b < (Max / a)) {
				return 0
			}
		}
	}
	return a * b
}

func division(a, b int) int {
	if a == Min && b == -1 {
		return 0
	}
	return a / b
}

func modulo(a, b int) int {
	if a == Min && b == -1 {
		return 0
	}
	return a % b
}
