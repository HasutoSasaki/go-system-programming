package signature

import (
	"encoding/binary"
	"fmt"
	"os"
)

// 簡単なテキストファイルを作成する関数
func CreateSimpleFile() {
	// 1. 普通のテキストファイル
	textFile, _ := os.Create("sample/signature-sample.txt")
	textFile.WriteString("This is a sample text file.\n")
	textFile.Close()

	// 2. 偽のPNGファイル（シグニチャだけ）
	fakePNG, _ := os.Create("sample/fake-signature.png")
	fakePNG.WriteString("\x89PNG\r\n\x1a\n")
	fakePNG.WriteString("これは偽のPNGファイルです。\n")
	fakePNG.Close()

	// 3. HTMLファイル
	htmlFile, _ := os.Create("sample/signature-sample.html")
	htmlFile.WriteString("<html><body><h1>This is a sample HTML file.</h1></body></html>\n")
	htmlFile.Close()

	fmt.Println("Sample files created successfully.")
	CheckFileSignature("sample/signature-sample.txt")
	CheckFileSignature("sample/fake-signature.png")
	CheckFileSignature("sample/signature-sample.html")
}

// 独自ファイル形式「MYDATA」を作成
func CreateCustomFile() {
	file, err := os.Create("sample/mydata.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 1. カスタムシグニチャ（４バイト)
	file.WriteString("MYDT")
	// 2. バージョン（１バイト）
	file.Write([]byte{0x01})
	// 3. データ長（4バイト、ビックエンディアン)
	data := "これは私のデータです"
	binary.Write(file, binary.BigEndian, uint32(len(data)))
	// 4. 実際のデータ
	file.WriteString(data)

	fmt.Println("Custom file created successfully.")
}
