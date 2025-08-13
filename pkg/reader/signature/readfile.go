package signature

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ReadCustomFile() {
	file, err := os.Open("sample/mydata.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 1. シグニチャ確認
	signature := make([]byte, 4)
	file.Read(signature)
	fmt.Printf("File Signature: % x\n", signature)

	if string(signature) != "MYDF" {
		fmt.Println("Invalid file signature")
		return
	}

	//2. バージョン読み取り
	version := make([]byte, 1)
	file.Read(version)
	fmt.Printf("File Version: % x\n", version)

	// 3. データ長読み取り
	var dataLength uint32
	binary.Read(file, binary.LittleEndian, &dataLength)
	fmt.Printf("Data Length: %d\n", dataLength)

	// 4. データ読み取り
	data := make([]byte, dataLength)
	file.Read(data)
	fmt.Printf("Data: % x\n", string(data))
}
