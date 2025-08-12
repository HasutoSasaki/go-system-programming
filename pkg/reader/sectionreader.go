package reader

import (
	"io"
	"os"
	"strings"
)

func SectionReader() {
	// SectionReader の使用例
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader) // 出力: "of io.Section"
}
