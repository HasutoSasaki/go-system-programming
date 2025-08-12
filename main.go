package main

import (
	"fmt"
	"go-system-programming/pkg/reader"
	"go-system-programming/pkg/writer"
)

func main() {
	fmt.Println("=== Writer Examples ===")

	fmt.Println("\n1. Basic File Writer:")
	writer.Basic()

	fmt.Println("\n2. Buffer Writer:")
	writer.BufferDemo()

	fmt.Println("\n3. Bufio Writer:")
	writer.BufioWriter()

	fmt.Println("\n3. Network TCP (commented out - requires internet):")
	writer.NetworkTCP()

	fmt.Println("\n4. HTTP Server (commented out - runs server):")
	writer.HTTPServer()

	fmt.Println("\n5. MultiWriter:")
	writer.MultiWriter()

	fmt.Println("\n6. Gzip Writer:")
	writer.GzipWriter()

	fmt.Println("\n7. Format Writer:")
	writer.FormatWriter()

	fmt.Println("\n8. JSON Writer:")
	writer.JSONWriter()

	fmt.Println("\n9. HTTP Request:")
	writer.HTTPRequest()

	fmt.Println("\n=== Reader Examples ===")

	fmt.Println("\n10. Section Reader:")
	reader.SectionReader()

	fmt.Println("\n11. File Reader (commented out - needs file.go):")
	reader.FileReader()

	fmt.Println("\n12. File Copy (commented out - needs sample.txt):")
	reader.FileCopy()

	fmt.Println("\n13. HTTP Response (commented out - requires internet):")
	reader.HTTPResponse()

	fmt.Println("\n14. Stdin Reader (commented out - interactive):")
	reader.StdinReader()
}
