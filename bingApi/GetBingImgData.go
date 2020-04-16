package GetBingImgData

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type BingData struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

const baseUrl = "https://cn.bing.com"

const src = baseUrl + "/HPImageArchive.aspx?format=js&n=10"

var JsonArr []string

var JsonStr string

//获取 bingImg
func GetData() {
	start()
}

func start() {
	data := httpGet(src)
	JsonArr = toJson(data)
	// JsonStr = json.Marshal(user)

	b, _ := json.Marshal(JsonArr)
	JsonStr = string(b)
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

func toJson(dataStr string) []string {
	var part BingData
	json.Unmarshal([]byte(dataStr), &part)
	var urlArr []string
	for _, val := range part.Images {
		var url = baseUrl + val.URL
		url = strings.Split(url, "&")[0]
		urlArr = append(urlArr, url)
	}
	return urlArr
}
