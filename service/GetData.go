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

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")

		fmt.Fprintf(w, cont)
	})
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		server.Shutdown(context.Background())
	})
	log.Println("启动服务 http://localhost:5000/getbingimg" + port)

	server.ListenAndServe() //设置监听的端口

}
