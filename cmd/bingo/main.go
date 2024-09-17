package main

/*
１〜５までの数字をランダムに100回出力し、「123」の並びが表れる度にビンゴ！と表示するプログラムを作成してください。

補足事項
- ビンゴ！と出力後に改行してください。
*/

import (
	"fmt"
	"math/rand"
)

const (
	bingoNum  string = "123"
	bingoWord string = "ビンゴ！"
	count     int    = 100
)

func BingoCheck(count int) {
	var res string

	for i := 0; i < count; i++ {
		randNum := rand.Intn(5) + 1
		res += fmt.Sprintf("%d", randNum)

		// 「123」が並んでいるかチェック
		if len(res) >= 3 && res[len(res)-3:] == bingoNum {
			res += bingoWord + "\n"
		}
	}

	fmt.Println(res)

}

func main() {
	BingoCheck(count)
	// 出力結果(例)：
	// 124535214123ビンゴ！
	// 34523123ビンゴ!
	// 3324...
}
