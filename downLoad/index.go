package downLoad

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

//获取 bingImg
func GetImg(url, filePath string) {
	var imgUrl = url

	res, err := http.Get(imgUrl)
	if err != nil {
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	//
	written, _ := io.Copy(writer, reader)
	// 输出文件字节大小
	fmt.Printf("Total length: %d", written)
}
