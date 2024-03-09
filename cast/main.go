package main

import "fmt"

func main() {
	var num1, num2 int

	// ユーザーから2つの整数値の入力を受け取る
	fmt.Print("1つ目の整数値を入力してください: ")
	fmt.Scanln(&num1)
	fmt.Print("2つ目の整数値を入力してください: ")
	fmt.Scanln(&num2)

	// 整数型から浮動小数点型へ変換する
	float1 := float64(num1)
	float2 := float64(num2)

	// 除算を行う
	result := float1 / float2

	// 結果を表示する
	fmt.Printf("%d ÷ %d = %.2f\n", num1, num2, result)
}
