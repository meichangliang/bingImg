package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DebugInfo struct {
	url string
}

func main() {
	data := httpGet("https://cn.bing.com/HPImageArchive.aspx?format=js&n=10")
	toJson(data)
}

func (dbgInfo DebugInfo) String() string {
	return fmt.Sprintf("{url: %s,}", dbgInfo.url)
}

func toJson(dataStr string) {
	data := dataStr

	var dbgInfos []DebugInfo
	json.Unmarshal([]byte(data), &dbgInfos)

	fmt.Println(dbgInfos)
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
