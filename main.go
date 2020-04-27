package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	bingApi "bingImg.com/bingApi"
	"bingImg.com/downLoad"
	service "bingImg.com/service"
	"github.com/robfig/cron"
)

var i = 0

const PATH = "./images"

func main() {
	// getImg()
	ticker()
	startService()
}

func ticker() {
	go func() {
		fmt.Println("启动定时任务")

		c := cron.New()
		spec := "0 0 1 * * ?" // 每天凌晨1点执行一次
		// spec := "*/5 * * * * ?" // 1秒钟来一次
		c.AddFunc(spec, func() {
			getImg()
		})
		c.Start()

		select {}
	}()
}

func getImg() {
	i++
	log.Println("cron running:", i)
	bingApi.GetData()
	fmt.Println("今日的bingImg", bingApi.JsonStr)

	for index, val := range bingApi.JsonArr {
		var strArr = strings.Split(val, ".")
		var lastName = strArr[len(strArr)-1]

		var imgPath = PATH + "/" + strconv.Itoa(index) + "." + lastName

		fmt.Println(val, imgPath)

		downLoad.GetImg(val, imgPath)
	}

	WriteToFile(PATH+"/data.json", bingApi.JsonStr)

}

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}

func startService() {

	const port = "5000"
	service.Start(port, PATH)

}
