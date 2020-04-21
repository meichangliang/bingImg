package GetData

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

func Start(cont string, port string, pathUrl string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getbingimg", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")

		fmt.Fprintf(w, cont)
	})
	mux.HandleFunc("/bz", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		idx := query["idx"]
		index := -1

		fileArr := getFileNum()

		if len(idx) > 0 {
			index, _ = strconv.Atoi(idx[0])
		}

		if index < 0 {
			rand.Seed(time.Now().Unix())
			index = rand.Intn(len(fileArr))
		} else {
			if index > len(fileArr) {
				index = len(fileArr) - 1
			}
			fmt.Println("第" + strconv.Itoa(index) + "张输出")
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "image/JPEG")
		fp := path.Join(pathUrl, fileArr[index])
		http.ServeFile(w, r, fp)

	})
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("closeServic关闭服务")
		server.Shutdown(context.Background())
	})
	log.Println("启动服务 http://localhost:5000/getbingimg")

	server.ListenAndServe() //设置监听的端口

}
func getFileNum() []string {
	files, _ := ioutil.ReadDir("./images")
	var nameArr []string
	for _, f := range files {
		nameArr = append(nameArr, f.Name())
	}
	return nameArr
}
