package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	// 出力先のzipファイルを作成
	file, err := os.Create("sample/output.zip")
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// zipファイル内にファイルを作成
	writer, err := zipWriter.Create("hello.txt")
	if err != nil {
		panic(err)
	}

	// string.Reader の内容をzipファイル内のファイルに書き込み
	reader := strings.NewReader("Hello, World!")
	_, err = io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}

	println("Created output.zip with hello.txt")
}
