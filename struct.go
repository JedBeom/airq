package airq

import (
	"time"
)

// AirQuality 구조체(보통 주석에서는 airq로 줄여부름)는 대기 오염 정보 필드들을 가지고 있는 구조체이다.
type AirQuality struct {
	DataTimeString string    `xml:"body>items>item>dataTime"` // 오염도 측정 연-월-일 시간:분
	DataTime       time.Time // DataTimeString 변수를 time.Time으로 파싱 후 저장하기 위함.

	MangName string `xml:"body>items>item>mangName"` // 측정망 정보 (국가배경, 교외대기, 도시대기, 도로변대기)

	So2Value float32 `xml:"body>items>item>so2Value"` // 아황산가스 농도 (ppm)
	Co2Value float32 `xml:"body>items>item>co2Value"` // 일산화탄소 농도 (ppm)
	O3Value  float32 `xml:"body>items>item>o3Value"`  // 오존 농도 (ppm)
	No2Value float32 `xml:"body>items>item>no2Value"` // 이산화질소 농도 (ppm)

	Pm10Value   int `xml:"body>items>item>pm10Value"`   // 미세먼지(PM10) 농도 (µg/m³)
	Pm10Value24 int `xml:"body>items>item>pm10Value24"` // 미세먼지(PM10) 24시간 예측 이동 농도 (µg/m³)
	Pm25Value   int `xml:"body>items>item>pm25Value"`   // 미세먼지(PM2.5) 농도 (µg/m³)
	Pm25Value24 int `xml:"body>items>item>pm25Value24"` // 미세먼지(body>items>item>pm2.5) 24시간 예측 이동 농도 (µg/m³)
	KhaiValue   int `xml:"body>items>item>khaiValue"`   // 통합대기환경수치

	KhaiGrade int `xml:"body>items>item>khaiGrade"` // 통합대기환경지수 (1-4)
	So2Grade  int `xml:"body>items>item>so2Grade"`  // 아황산가스 지수 (1-4)
	CoGrade   int `xml:"body>items>item>coGrade"`   // 일산화탄소 지수 (1-4)
	O3Grade   int `xml:"body>items>item>o3Grade"`   // 오존 지수 (1-4)
	No2Grade  int `xml:"body>items>item>no2Grade"`  // 이산화질소 지수 (1-4)

	Pm10GradeKorea   int `xml:"body>items>item>pm10Grade1h"` // 미세먼지(PM10) 1시간 등급자료 한국 기준 (1-4)
	Pm10Grade24Korea int `xml:"body>items>item>pm10Grade"`   // 미세먼지(PM10) 24시간 등급자료 한국 기준 (1-4)
	Pm10GradeWHO     int // 미세먼지(PM10) 1시간 등급자료 WHO 기준 (1-8)
	Pm10Grade24WHO   int // 미세먼지(PM10) 24시간 등급자료 WHO 기준 (1-8)

	Pm25GradeKorea   int `xml:"body>items>item>pm25Grade1h"` // 미세먼지(PM2.5) 1시간 등급 자료 한국 기준 (1-4)
	Pm25Grade24Korea int `xml:"body>items>item>pm25Grade"`   // 미세먼지(PM2.5) 24시간 등급 자료 한국 기준 (1-4)
	Pm25GradeWHO     int // 미세먼지(PM2.5) 1시간 등급자료 WHO 기준 (1-8)
	Pm25Grade24WHO   int // 미세먼지(PM2.5) 24시간 등급자료 WHO 기준 (1-8)
}

// apiError 구조체는 api 이용 시 에러가 났을 경우에 그 에러 내용을 파싱하기 위한 구조체이다.
// error가 아니다.
type apiError struct {
	Code string `xml:"header>resultCode"`
	Msg  string `xml:"header>resultMsg"`
}
