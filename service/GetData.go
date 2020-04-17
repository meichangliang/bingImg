package GetData

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(cont string, port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getbingimg", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, cont) //这个写入到w的是输出到客户端的
	})
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		server.Shutdown(context.Background())
	})
	log.Println("启动服务" + port)

	server.ListenAndServe() //设置监听的端口

}
