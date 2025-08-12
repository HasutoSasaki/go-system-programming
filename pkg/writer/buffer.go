package writer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func BufferDemo() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	buffer.WriteString("bytes.Buffer example\n")
	io.WriteString(&buffer, "bytes.Buffer example\n")
	fmt.Println(buffer.String())
}

func BufioWriter() {
	// buffer writer
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
