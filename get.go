package airq

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
)

var (
	apiLink = "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getMsrstnAcctoRltmMesureDnsty?stationName=%v&dataTerm=DAILY&pageNo=1&numOfRows=%v&ServiceKey=%v&ver=1.3"
)

// ByStation 함수는 stationName을 기반으로 Airq를 가져온다. 몇개의 정보를 가져올 지 정할 수 있다.
func ByStation(stationName string, rows int) (qualities []AirQuality, err error) {

	if serviceKey == "" {
		err = errors.New("Service Key wasn't imported.")
		err = errors.Wrap(err, "airq")
		return
	}

	link := fmt.Sprintf(apiLink, stationName, rows, serviceKey) // 링크 생성

	resp, err := http.Get(link) // Get으로 응답을 받는다.
	defer func() {
		_ = resp.Body.Close() // 함수 종료 시 resp을 Close
	}()
	if err != nil {
		return
	}

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
	defer func() {
		time.Local = originalLocal // time.Local을 원래대로 변경.
	}()

	for i := range qualities {
		qualities[i].StationName = stationName

		// Pm10GradeWHO, Pm25GradeWHO에 WHO 기준 지수 대입
		qualities[i].Pm10GradeWHO, qualities[i].Pm25GradeWHO = whoGradeRater(qualities[i].Pm10Value, qualities[i].Pm25Value)
		// Pm10Grade24WHO, Pm25Grade24WHO에 WHO 기준 지수 대입
		qualities[i].Pm10Grade24WHO, qualities[i].Pm25Grade24WHO = whoGradeRater(qualities[i].Pm10Value24, qualities[i].Pm25Value24)
		// DataTimeString을 time.Time으로 파싱.

		qualities[i].DataTimeString = strings.Replace(qualities[i].DataTimeString, "24:", "00:", 1)
		qualities[i].DataTime, err = dateparse.ParseLocal(qualities[i].DataTimeString)
		if err != nil {
			return
		}
	}

	return

}

// NowByStation 함수는 ByStation 함수를 이용해 현재의 Airq만 리턴한다.
func NowByStation(stationName string) (quality AirQuality, err error) {
	qualities, err := ByStation(stationName, 1)
	if err != nil {
		return
	}

	if len(qualities) == 0 {
		err = errors.New("No airq;")
		err = errors.Wrap(err, "airq")
		return
	}
	quality = qualities[0]
	return
}
