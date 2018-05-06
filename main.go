package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Println("aiueo")
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
