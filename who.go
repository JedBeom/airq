package airq

func whoGradeRater(pm10 int, pm25 int) (pm10GradeWHO int, pm25GradeWHO int) {
	// 미세먼지 등급을 메기기(1-8)
	// 등급이 낮을 수록 공기의 상태가 좋음
	switch {
	case pm10 <= 15:
		pm10GradeWHO = 1
	case pm10 <= 30:
		pm10GradeWHO = 2
	case pm10 <= 40:
		pm10GradeWHO = 3
	case pm10 <= 50:
		pm10GradeWHO = 4
	case pm10 <= 75:
		pm10GradeWHO = 5
	case pm10 <= 100:
		pm10GradeWHO = 6
	case pm10 <= 150:
		pm10GradeWHO = 7
	default:
		pm10GradeWHO = 8
	}
	// 초미세먼지 등급을 메김(1-8)
	// 등급이 낮을수록 공기의 상태가 좋음
	switch {
	case pm25 <= 8:
		pm25GradeWHO = 1
	case pm25 <= 15:
		pm25GradeWHO = 2
	case pm25 <= 20:
		pm25GradeWHO = 3
	case pm25 <= 25:
		pm25GradeWHO = 4
	case pm25 <= 37:
		pm25GradeWHO = 5
	case pm25 <= 50:
		pm25GradeWHO = 6
	case pm25 <= 75:
		pm25GradeWHO = 7
	default:
		pm25GradeWHO = 8
	}
	return
}
