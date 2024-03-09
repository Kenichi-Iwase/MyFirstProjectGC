package main

import "fmt"

func main() {
	var number int

	// ユーザーから数値の入力を受け取る
	fmt.Print("数値を入力してください: ")
	fmt.Scanln(&number)

	// for文を使って、1から10までの数値を掛ける
	for i := 1; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d × %d = %d\n", number, i, result)
	}
}
