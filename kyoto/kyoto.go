package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// レスポンスの形式に合わせて struct を定義する
type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func GetWeather() string {
	jsonStr := httpGetStr("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/260000.json")
	weather := formatWeather(jsonStr)

	area := fmt.Sprintf("%sの天気です。\n", weather.Area)
	head := fmt.Sprintf("%s\n", weather.HeadLine)
	body := fmt.Sprintf("%s\n", weather.Body)
	result := area + head + body

	return result
}

func httpGetStr(url string) string {
	// HTTPリクエストを発行しレスポンスを取得する
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	return string(body)
}

func formatWeather(str string) *Weather {
	weather := new(Weather)
	if err := json.Unmarshal([]byte(str), weather); err != nil {
		log.Fatal("JSON Unmarshal error:", err)
	}
	return weather
}
