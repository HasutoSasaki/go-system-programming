package writer

import "os"

func StdoutDemo() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
