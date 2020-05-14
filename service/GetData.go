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
	"strings"
	"time"
)

func Start(port string, pathUrl string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/getbingimg", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")

		data, err := ioutil.ReadFile(pathUrl + "/data.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, string(data))
	})
	mux.HandleFunc("/bz", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		idx := query["idx"]
		index := -1

		fileArr := getFileNum(pathUrl)

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
		}

		if index > len(fileArr)-1 {
			index = len(fileArr) - 1
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

func getFileNum(pathUrl string) []string {
	files, _ := ioutil.ReadDir(pathUrl)
	var nameArr []string
	for _, f := range files {
		countSplit := strings.Split(f.Name(), ".")
		if countSplit[1] != "json" {
			nameArr = append(nameArr, f.Name())
		}
	}
	return nameArr
}
