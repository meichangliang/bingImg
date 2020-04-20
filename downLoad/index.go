package downLoad

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var imgPath = "/"

//获取 bingImg
func GetImg(url, filePath string) {
	var imgUrl = url

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	imgPath := "./"
	// 根据图片url获取其文件名
	fileName := path.Base(imgUrl)

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
}
