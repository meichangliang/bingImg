package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	bingApi "bingImg.com/bingApi"
	"bingImg.com/downLoad"
	service "bingImg.com/service"
	"github.com/robfig/cron"
)

var num int

var i = 0

func main() {
	start()

	c := cron.New()
	// spec := "0 0 1 * * ?" // 每天凌晨1点执行一次
	spec := "*/1 * * * * ?" // 1秒钟来一次
	c.AddFunc(spec, func() {
		start()
	})
	c.Start()

	select {}

}

func start() {
	i++
	log.Println("cron running:", i)
	startService()
}

const PATH = "./images"

func startService() {
	bingApi.GetData()

	for index, val := range bingApi.JsonArr {
		var strArr = strings.Split(val, ".")
		var lastName = strArr[len(strArr)-1]

		var imgPath = PATH + "/" + strconv.Itoa(index) + "." + lastName

		fmt.Println(val, imgPath)

		downLoad.GetImg(val, imgPath)
	}

	const port = "5000"
	num++
	fmt.Println("启动次数", num)
	http.Get("http://localhost:" + port + "/close")
	start()

	service.Start(bingApi.JsonStr, port, PATH)

}
