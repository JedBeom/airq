package airq

import (
	"encoding/xml"
	"time"
)

type AirQuality struct {
	DataTime time.Time `xml:"dataTime"` // 오염도 측정 연-월-일 시간:분

	So2Value float32 `xml:"so2Value"` // 아황산가스 농도 (ppm)
	Co2Value float32 `xml:"co2Value"` // 일산화탄소 농도 (ppm)
	O3Value  float32 `xml:"o3Value"`  // 오존 농도 (ppm)
	No2Value float32 `xml:"no2Value"` // 이산화질소 농도 (ppm)

	Pm10Value   int `xml:"pm10Value"`   // 미세먼지(PM10) 농도 (µg/m³)
	Pm10Value24 int `xml:"pm10Value24"` // 미세먼지(PM10) 24시간 예측 이동 농도 (µg/m³)
	Pm25Value   int `xml:"pm25Value"`   // 미세먼지(PM2.5) 농도 (µg/m³)
	Pm25Value24 int `xml:"pm25Value24"` // 미세먼지(PM2.5) 24시간 예측 이동 농도 (µg/m³)
	KhaiValue   int `xml:"khaiValue"`   // 통합대기환경수치

	KhaiGrade int `xml:"khaiGrade"` // 통합대기환경지수 (1-4)
	So2Grade  int `xml:"so2Grade"`  // 아황산가스 지수 (1-4)
	CoGrade   int `xml:"coGrade"`   // 일산화탄소 지수 (1-4)
	O3Grade   int `xml:"o3Grade"`   // 오존 지수 (1-4)
	No2Grade  int `xml:"no2Grade"`  // 이산화질소 지수 (1-4)

	Pm10GradeKorea   int `xml:"pm10Grade1h"` // 미세먼지(PM10) 1시간 등급자료 한국 기준 (1-4)
	Pm10Grade24Korea int `xml:"pm10Grade"`   // 미세먼지(PM10) 24시간 등급자료 한국 기준 (1-4)
	Pm10GradeWHO     int // 미세먼지(PM10) 1시간 등급자료 WHO 기준 (1-8)
	Pm10Grade24WHO   int // 미세먼지(PM10) 24시간 등급자료 WHO 기준 (1-8)

	Pm25GradeKorea   int `xml:"pm25Grade1h"` // 미세먼지(PM2.5) 1시간 등급 자료 한국 기준 (1-4)
	Pm25Grade24Korea int `xml:"pm25Grade"`   // 미세먼지(PM2.5) 24시간 등급 자료 한국 기준 (1-4)
	Pm25GradeWHO     int // 미세먼지(PM2.5) 1시간 등급자료 WHO 기준 (1-8)
	Pm25Grade24WHO   int // 미세먼지(PM2.5) 24시간 등급자료 WHO 기준 (1-8)
}

type body struct {
	airQuality AirQuality `xml:"body>items>item"`
}
