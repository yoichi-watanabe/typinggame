package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var correct int

func input(r io.Reader) {

	go func() {
		s := bufio.NewScanner(r)
		var question string
		var answer string

		for s.Scan() {

			// 正解数をインクリメント
			answer = s.Text()
			if question == answer {
				correct++
			}
			question = output()

			fmt.Println(question)
			fmt.Print(">")
		}
	}()
}

// 標準出力に英単語を出す
func output() string {

	// 読み込み用にファイルを開く
	dictCsv, err := os.Open("./dict.csv")
	if err != nil {
		//return err
		fmt.Println(err)
	}
	// 関数終了時に閉じる
	defer dictCsv.Close()

	// csv英単語辞書データを読み込み
	r := csv.NewReader(dictCsv)
	rows, err := r.ReadAll()

	// 乱数を設定し、辞書データ内のランダムな値を返却する
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(len(rows[0]))
	return rows[0][randomInt]
}

func main() {

	correct = -1
	fmt.Println("制限時間は30秒です。Enter を押すと開始します。")

	input(os.Stdin)

	time.Sleep(time.Second * 30)

	fmt.Println("★★★TIME UP★★★")
	fmt.Print("合計性回数 : ")
	fmt.Println(correct)
}
