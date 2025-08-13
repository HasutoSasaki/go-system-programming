package reader

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
)

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)
	// CRCを計算して追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func EmbedTextPNG() {
	file, err := os.Open("sample/Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("sample/Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	// シグニチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	dumpChunk(chunks[0])
	io.Copy(newFile, chunks[0])
	// テキストチャンクを追加
	textChunkReader := textChunk("ASCII PROGRAMMING++")
	dumpChunk(textChunkReader)
	io.Copy(newFile, textChunkReader)

	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		dumpChunk(chunk)
		io.Copy(newFile, chunk)
	}
}
