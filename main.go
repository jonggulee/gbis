package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jonggulee/gbis/utils"
)

const (
	busArrivalAPIAddress = "http://apis.data.go.kr/6410000/busarrivalservice/getBusArrivalList"
	stationID            = "204000235"
)

type BusArrivalList struct {
	Flag           string `xml:"flag"`
	LocationNo1    int    `xml:"locationNo1"`
	LocationNo2    int    `xml:"locationNo2"`
	LowPlate1      int    `xml:"lowPlate1"`
	LowPlate2      int    `xml:"lowPlate2"`
	PlateNo1       string `xml:"plateNo1"`
	PlateNo2       string `xml:"plateNo2"`
	PredictTime1   int    `xml:"predictTime1"`
	PredictTime2   int    `xml:"predictTime2"`
	RemainSeatCnt1 int    `xml:"remainSeatCnt1"`
	RemainSeatCnt2 int    `xml:"remainSeatCnt2"`
	RouteID        int    `xml:"routeId"`
	StaOrder       int    `xml:"staOrder"`
	StationID      int    `xml:"stationId"`
}

type Response struct {
	BusArrivalList []BusArrivalList `xml:"msgBody>busArrivalList"`
}

func main() {
	serviceKey := os.Getenv("serviceKey")
	if serviceKey == "" {
		fmt.Println("service key 확인")
	}

	url := fmt.Sprintf("%s?serviceKey=%s&stationId=%s", busArrivalAPIAddress, serviceKey, stationID)
	fmt.Println(url)

	resp, err := http.Get(url)
	utils.HandleErr(err)

	defer resp.Body.Close()

	// 응답 바디 읽기
	xmlData, err := io.ReadAll(resp.Body)
	utils.HandleErr(err)

	var response Response
	err = xml.Unmarshal(xmlData, &response)
	utils.HandleErr(err)

	for _, busArrival := range response.BusArrivalList {
		fmt.Printf("%d분후 도착예정 / ", busArrival.PredictTime1)
		fmt.Printf("%d번째전 정류소\n", busArrival.LocationNo1)
		fmt.Println("")
	}
}
