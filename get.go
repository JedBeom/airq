package airq

import (
	"encoding/xml"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	byStationLink = "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getMsrstnAcctoRltmMesureDnsty?stationName=%v&dataTerm=DAILY&pageNo=1&numOfRows=%v&ServiceKey=%v&ver=1.3"
)

// GetAirqByStation 함수는 stationName을 기반으로 Airq를 가져온다. 몇개의 정보를 가져올 지 정할 수 있다.
func GetAirqByStation(stationName string, rows int) (qualities []AirQuality, err error) {
	link := fmt.Sprintf(byStationLink, stationName, rows, serviceKey) // 링크 생성

	resp, err := http.Get(link) // Get으로 응답을 받는다.
	if err != nil {
		return
	}
	defer resp.Body.Close() // 함수 종료 시 resp을 Close

	data, err := ioutil.ReadAll(resp.Body) // resp.Body를 ioutil로 변환.
	if err != nil {
		return
	}

	var aE apiError                // 만약 에러가 났다면, 그 정보를 파싱하기 위한 변수.
	err = xml.Unmarshal(data, &aE) // 파싱 시도.
	if err != nil {
		return
	}

	if aE.Msg != "NORMAL SERVICE." { // 만약 정상적인 응답이 아니라면
		err = errors.New(aE.Msg)       // resultMsg의 내용을 에러로 변환.
		err = errors.Wrap(err, "airq") // 에러 앞에 airq: 를 덧붙임.
		return
	}

	err = xml.Unmarshal(data, &qualities) // XML 파싱

	originalLocal := time.Local                 // 원래의 time.Local 로드.
	loc, err := time.LoadLocation("Asia/Seoul") // Asia/Seoul(KST)를 time.Location으로 변환.
	if err != nil {
		return
	}
	time.Local = loc // Asia/Seoul을 이 기기의 Local로 변환.

	for i := range qualities {
		// Pm10GradeWHO, Pm25GradeWHO에 WHO 기준 지수 대입
		qualities[i].Pm10GradeWHO, qualities[i].Pm25GradeWHO = whoGradeRater(qualities[i].Pm10Value, qualities[i].Pm25Value)
		// Pm10Grade24WHO, Pm25Grade24WHO에 WHO 기준 지수 대입
		qualities[i].Pm10Grade24WHO, qualities[i].Pm25Grade24WHO = whoGradeRater(qualities[i].Pm10Value24, qualities[i].Pm25Value24)
		// DataTimeString을 time.Time으로 파싱.

		qualities[i].DataTimeString = strings.Replace(qualities[i].DataTimeString, "24:", "00:", 1)
		qualities[i].DataTime, err = dateparse.ParseLocal(qualities[i].DataTimeString)
		if err != nil {
			time.Local = originalLocal // time.Local을 원래대로 변경
			return
		}
	}

	time.Local = originalLocal // time.Local을 원래대로 변경.
	return

}

// GetAirqOfNowByStation 함수는 GetAirqByStation 함수를 이용해 현재의 Airq만 리턴한다.
func GetAirqOfNowByStation(stationName string) (quality AirQuality, err error) {
	qualities, err := GetAirqByStation(stationName, 1)
	if err != nil {
		return
	}

	if len(qualities) == 0 {
		err = errors.New("No airq; Maybe stationName was wrong")
		err = errors.Wrap(err, "airq")
		return
	}
	quality = qualities[0]
	return
}
