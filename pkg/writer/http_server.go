package writer

import (
	"net/http"
)

func HTTPServer() {
	// http.ResponseWriter
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http.ResponseWriter sample!\n"))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
