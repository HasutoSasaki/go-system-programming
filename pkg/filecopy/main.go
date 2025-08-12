package main

import (
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <src> <dst>\nExample: %s old.txt new.txt\n", os.Args[0], os.Args[0])
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		usage()
		os.Exit(2)
	}
	srcPath := args[0]
	dstPath := args[1]

	src, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		panic(err)
	}
}
