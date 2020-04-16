package getData

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func Start() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.0:8000", nil)
}
