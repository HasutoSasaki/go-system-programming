package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func FormatWriter() {
	// フォーマットしてデータをio.Writerに書き出す
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}

func JSONWriter() {
	// JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
