package writer

import (
	"compress/gzip"
	"os"
)

func GzipWriter() {
	// gzip
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	writer.Write([]byte("gzip example\n"))
	writer.Close()
}
