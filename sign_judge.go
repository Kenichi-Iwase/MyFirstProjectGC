package main

import "fmt"

func main() {
	var number int

	// ユーザーから数値の入力を受け取る
	fmt.Print("数値を入力してください: ")
	fmt.Scanln(&number)

	// 条件分岐を使って、数値の正負を判定する
	if number > 0 {
		fmt.Printf("%d は正の数です。\n", number)
	} else if number < 0 {
		fmt.Printf("%d は負の数です。\n", number)
	} else {
		fmt.Printf("%d はゼロです。\n", number)
	}
}
