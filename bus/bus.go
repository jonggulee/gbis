package bus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	busArrivalAPIAddress          = "http://ws.bus.go.kr/api/rest/stationinfo/getStationByUid"
	WiryeJungangMiddleSchoolArsId = "48626" // 위례중앙중학교 정류소 번호
	WiryeStarfieldCityArsId       = "28532" // 하남시위례도서관.스타필드시티.위례지웰푸르지오
	resultType                    = "json"
)

type JsonResponse struct {
	MsgBody MsgBody `json:"msgBody"`
}

type MsgBody struct {
	Bus []Bus `json:"itemList"`
}

type Bus struct {
	StNm         string `json:"stNm"`         // 정류소명
	BusRouteAbrv string `json:"busRouteAbrv"` // 도착번호
	Arrmsg1      string `json:"arrmsg1"`      // 첫번째 도착예정 버스의 도착정보 메시지
	Arrmsg2      string `json:"arrmsg2"`      // 두번째 도착예정 버스의 도착정보 메시지
	StationNm1   string `json:"stationNm1"`   // 첫번째 도착예정 버스의 현재 정류소명
	StationNm2   string `json:"stationNm2"`   // 두번째 도착예정 버스의 현재 정류소명
}

func CheckServiceKey() (string, error) {
	serviceKey := os.Getenv("serviceKey")
	if serviceKey == "" {
		err := errors.New("service key 확인")
		return "", err
	}
	return serviceKey, nil
}

func GetArrivalBus(arsId string) []Bus {
	serviceKey, err := CheckServiceKey()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	url := fmt.Sprintf("%s?serviceKey=%s&arsId=%s&resultType=%s", busArrivalAPIAddress, serviceKey, arsId, resultType)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var data JsonResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var filteredBuses []Bus
	if arsId == WiryeJungangMiddleSchoolArsId {
		for _, bus := range data.MsgBody.Bus {
			if bus.BusRouteAbrv == "333" || bus.BusRouteAbrv == "440" || bus.BusRouteAbrv == "315" {
				filteredBuses = append(filteredBuses, bus)
			}
		}
	} else if arsId == WiryeStarfieldCityArsId {
		for _, bus := range data.MsgBody.Bus {
			if bus.BusRouteAbrv == "3217" || bus.BusRouteAbrv == "231" {
				filteredBuses = append(filteredBuses, bus)
			}
		}
	}

	return filteredBuses
}
