package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.zip")

	//zipファイルを直接レスポンスに書き込み
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// zipファイル内にファイルを作成
	fileWriter, err := zipWriter.Create("sample.txt")
	if err != nil {
		http.Error(w, "Failed to create zip file", http.StatusInternalServerError)
		return
	}

	// 文字列をzipファイルに書き込む
	reader := strings.NewReader("Hello, this is a sample text.")
	io.Copy(fileWriter, reader)
}

func main() {
	http.HandleFunc("/download", handler)
	http.ListenAndServe(":8080", nil)
	println("Server started at http://localhost:8080/download")
}
