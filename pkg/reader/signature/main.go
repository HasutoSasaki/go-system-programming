package signature

import "fmt"

func Demo() {
	fmt.Println("=== Signature Demo ===")

	// 1. 各種ファイル作成
	CreateSimpleFile()

	// 2. カスタムファイル作成・読み取り
	CreateCustomFile()
	ReadCustomFile()

	// 3. 既存ファイルがあれば確認
	CheckFileSignature("sample/Lenna.png")
}
