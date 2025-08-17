package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		file, _ := os.Create(fmt.Sprintf("test%d.txt", i))
		file.WriteString("Hello World")
		file.Close()
		time.Sleep(1 * time.Second)
	}
}
