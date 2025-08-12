package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// 適当なサイズのランダムなファイルを生成するプログラム
func main() {
	// 1024バイトのランダムなファイルを作成
	file, err := os.Create("sample/random.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// io.CopyN で指定したバイト数だけコピー
	n, err := io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated random file with %d bytes\n", n)
}

// 出力後のバイナリファイルを見る方法
// hexdump -C sample/random.bin | head
