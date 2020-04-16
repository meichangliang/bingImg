package main

import (
	bingApi "bingImg.com/bingApi"
	service "bingImg.com/service"
)

func main() {
	bingApi.GetData()
	// var data = bingApi.JsonArr

	service.Start()
}
