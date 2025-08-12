package writer

import (
	"io"
	"net"
	"os"
)

func NetworkTCP() {
	// internet access
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	req := "GET / HTTP/1.1\r\n" +
		"Host: ascii.jp\r\n" +
		"User-Agent: go-net-dial/1.0\r\n" +
		"Connection: close\r\n\r\n"
	conn.Write([]byte(req))
	io.Copy(os.Stdout, conn)
}
