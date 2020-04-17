package main

import (
	bingApi "bingImg.com/bingApi"
	service "bingImg.com/service"
)

func main() {
	startService()
}

func startService() {
	bingApi.GetData()
	service.Start(bingApi.JsonStr, "5000")
}
