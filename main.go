package main

import (
	"fmt"
	"log"
	"net/http"

	bingApi "bingImg.com/bingApi"
	service "bingImg.com/service"
	"github.com/robfig/cron"
)

var num int

func main() {
	i := 0
	c := cron.New()
	spec := "0 0 1 * * ?" // 每天凌晨1点执行一次
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		startService()
	})
	c.Start()

	select {}

}

func startService() {
	bingApi.GetData()
	const port = "5000"
	num++
	fmt.Println("启动次数", num)
	http.Get("http://localhost:" + port + "/getbingimg")
	service.Start(bingApi.JsonStr, "5000")
}
