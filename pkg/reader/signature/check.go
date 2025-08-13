package signature

import (
	"fmt"
	"os"
)

func CheckFileSignature(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 先頭16バイトを読み取り
	signature := make([]byte, 16)
	n, err := file.Read(signature)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("File: %s\n", filename)
	fmt.Printf("先頭%dバイト（16進数）: % x\n", n, signature[:n])
	fmt.Printf("先頭%dバイト（文字列）: %q\n", n, string(signature[:n]))

	// ファイル形式を判定
	identifyFileType(signature[:n])
	fmt.Println("---")
}

func identifyFileType(signature []byte) {
	if len(signature) < 4 {
		fmt.Println("ファイル形式：不明(データが少なすぎます)")
		return
	}

	switch {
	case signature[0] == 0x89 && string(signature[1:4]) == "PNG":
		fmt.Println("ファイル形式：PNG画像")
	case signature[0] == 0xFF && signature[1] == 0xD8:
		fmt.Println("ファイル形式：JPEG画像")
	case string(signature[0:4]) == "GIF8":
		fmt.Println("ファイル形式：GIF画像")
	case signature[0] == 0x50 && signature[1] == 0x4B:
		fmt.Println("ファイル形式：ZIP/JAR/DOCX等")
	case string(signature[0:4]) == "%PDF":
		fmt.Println("ファイル形式：PDF文書")
	case string(signature[0:4]) == "<!DO" || string(signature[0:4]) == "<htm":
		fmt.Println("ファイル形式: HTML")
	default:
		fmt.Println("ファイル形式：不明")
	}
}
