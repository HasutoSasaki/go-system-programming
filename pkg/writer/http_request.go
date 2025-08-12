package writer

import (
	"net/http"
	"os"
)

func HTTPRequest() {
	// HTTPリクエストを作成して、ヘッダーを追加して出力する
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できる")
	request.Write(os.Stdout)
}
