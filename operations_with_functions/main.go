package main

import "fmt"

// 2つの整数値の合計を計算する関数
func add(a, b int) int {
	return a + b
}

// 2つの整数値の差を計算する関数
func subtract(a, b int) int {
	return a - b
}

// 2つの整数値の積を計算する関数
func multiply(a, b int) int {
	return a * b
}

// 2つの整数値の商を計算する関数
func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var num1, num2 int

	// ユーザーから2つの整数値の入力を受け取る
	fmt.Print("1つ目の整数値を入力してください: ")
	fmt.Scanln(&num1)
	fmt.Print("2つ目の整数値を入力してください: ")
	fmt.Scanln(&num2)

	// 関数を呼び出して計算を行う
	sum := add(num1, num2)
	diff := subtract(num1, num2)
	product := multiply(num1, num2)
	quotient := divide(num1, num2)

	// 結果を表示する
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	fmt.Printf("%d - %d = %d\n", num1, num2, diff)
	fmt.Printf("%d × %d = %d\n", num1, num2, product)
	fmt.Printf("%d ÷ %d = %.2f\n", num1, num2, quotient)
}
