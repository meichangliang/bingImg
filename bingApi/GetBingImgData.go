package GetBingImgData

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type BingData struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

const baseUrl = "https://cn.bing.com"

var JsonArr []string

var JsonStr string

//获取 bingImg
func GetData() {
	start()
}

func start() {

	var ImgUrlArr []string

	for i := 0; i < 8; i++ {
		var src = baseUrl + "/HPImageArchive.aspx?format=js&n=1&idx=" + strconv.Itoa(i)
		data := httpGet(src)
		JsonArr = toJson(data)
		ImgUrlArr = append(ImgUrlArr, JsonArr[0])
	}

	JsonArr = ImgUrlArr

	b, _ := json.Marshal(ImgUrlArr)
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
