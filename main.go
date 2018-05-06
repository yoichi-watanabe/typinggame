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

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(ch)
	}()
	// TODO: チャネルを返す
	return ch
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
	ch := input(os.Stdin)
	var correct int

	for {
		question := output()
		fmt.Println(question)
		fmt.Print(">")

		// 正解数をインクリメント
		if question == <-ch {
			correct++
			fmt.Println(correct)
		}
	}
}
