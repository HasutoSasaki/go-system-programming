package writer

import (
	"io"
	"os"
)

func MultiWriter() {
	// multi writer
	file, err := os.Create("multi_writer.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter sample!\n")
}
