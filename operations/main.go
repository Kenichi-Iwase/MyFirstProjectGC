package main

import "fmt"

func main() {
	// 変数の定義
	a := 10
	b := 5

	// 加算
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// 減算
	diff := a - b
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	// 乗算
	product := a * b
	fmt.Printf("%d * %d = %d\n", a, b, product)

	// 除算
	quotient := a / b
	fmt.Printf("%d / %d = %d\n", a, b, quotient)

	// 剰余
	remainder := a % b
	fmt.Printf("%d %% %d = %d\n", a, b, remainder)
}
