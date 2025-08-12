package reader

import (
	"fmt"
	"io"
	"os"
)

func StdinReader() {
	// 標準入力
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

func BasicRead() {
	// 基本的な読み込み例
	// 1024バイトのバッファをmakeで作る
	// buffer := make([]byte, 1024)
	// sizeは実際に読み込んだバイト数、errはエラー
	// size, err := r.Read(buffer)
}

func ReadHelpers() {
	// 読み込みの補助関数
	// 全て読み込む
	// buffer, err := io.ReadAll(reader)
	// 決まったバイト数だけ確実に読む
	// buffer := make([]byte, 4)
	// size, err := io.ReadFull(reader, buffer)
}

func CopyHelpers() {
	// コピーの補助関数
	// 全てコピー
	// writeSize, err := io.Copy(writer, reader)
	// 指定したサイズだけ
	// writeSize, err := io.CopyN(writer, reader, size)
	// あらかじめコピーする量が決まっていて、無駄なバッファを使いたくない場合
	//8KBのバッファを使う
	// buffer := make([]byte, 8*1024)
	// io.CopyBuffer(writer, reader, buffer)
}
