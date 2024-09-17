package main

import "fmt"

/*
for文を使わずに、beginからendまでの数についてFizzBuzz問題を解く関数FizzBuzz()を実装してください。

補足事項
- main関数に手を加えずに実行できるようにしてください。
- 必要に応じて新しい関数を定義しても構いません。

※FizzBuzz問題とは
3の倍数のときはFizz、5の倍数のときはBuzz、3と5両方の倍数の場合にはFizzBuzzと出力し、それ以外の場合はその数を出力するという問題です。
*/

const (
	begin int64 = 1
	end   int64 = 100
)

func FizzBuzz(begin, end int64) {
	// endに達したら早期リターン
	if begin > end {
		return
	}

	var res string
	switch {
	case begin%3 == 0:
		res = "Fizz"
	case begin%5 == 0:
		res = "Buzz"
	case begin%3 == 0 && begin%5 == 0:
		res = "FizzBuzz"
	default:
		res = fmt.Sprintf("%d", begin)
	}
	fmt.Println(res)

	// 再帰呼び出し
	FizzBuzz(begin+1, end)

}

// mainは手を加えない
func main() {
	FizzBuzz(begin, end)
	// 出力結果
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// (中略)
	// 14
	// FizzBuzz
	// 16
	// (中略)
	// 98
	// Fizz
	// Buzz
}
