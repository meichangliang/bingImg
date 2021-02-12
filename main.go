package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	bingApi "bingImg.com/bingApi"
	"bingImg.com/downLoad"
	service "bingImg.com/service"
	"github.com/robfig/cron"
	utils "bingImg.com/utils"
)

var i = 0

const PATH = "./images"

func main() {
	getImg()
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

	utils.WriteToFile(PATH+"/data.json", bingApi.JsonStr)

}


func startService() {

	const port = "5000"
	service.Start(port, PATH)

}
