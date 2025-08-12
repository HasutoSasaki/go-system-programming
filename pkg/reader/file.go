package reader

import (
	"io"
	"os"
)

func FileReader() {
	// ファイル読み込み
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close() // 現在のスコープでの処理が終わったらファイルを閉じる
	io.Copy(os.Stdout, file)
}

func FileCopy() {
	// 練習：fileのコピー
	srcPath := "sample.txt"      // コピー元ファイル
	dstPath := "sample_copy.txt" // コピー先ファイル

	src, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		panic(err)
	}
}
