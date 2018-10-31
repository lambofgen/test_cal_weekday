package services

import (
	"errors"
	"math"
)

var (
	// ปีที่น้อยที่สุด
	baseYear float64 = 1900
	// สร้าง slice ของวันขึ้นมาโดย เริ่มจาก Monday เพราะโจทย์กำหนดว่า วันที่น้อยสุดคือ 1 jan 1900 ซึ่งเป็นวันจันทร์  (ใช้วันจันทร์เป็นหลัก)
	weekDay = []string{"Monday", "Tuesday", "Wedsday", "Thursday", "Friday", "Saturday", "Sunday"}
	// สร้าง slice ของวันจำนวนวันทั้งหมดของเดือน โดยใช้ index + 1 แทนหมายเลขเดือน (เดือน 2 ใส่ 0 เพราะจะไปคำนวณต่างหาก)
	dayOfMonth = []int{31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	errDateOutOfLenght = errors.New("Date input out of length")
)

func CalWeekDay(targetYear float64, targetMonth float64, targetDay float64) (string, error) {
	// เช็ควันที่ว่าตรงกับความจริงไหม
	if isDateOutOfLenght(targetYear, targetMonth, targetDay) {
		return "", errDateOutOfLenght
	}
	var sumDayPastYear float64
	// ถ้า baseYear == targetYear ก็ไม่จำเป็นต้องคำนวณจำนวนวันของปีที่ผ่านมา
	if baseYear != targetYear {
		// คำนวณหาจำนวน leap year
		countLeapYear, err := calCountOfLeapYear(baseYear, targetYear-1)
		if err != nil {
			return "", err
		}
		// หาผลต่างระหว่างปี base กับ ปีก่อนปีเป้าหมาย เพื่อนำไปคำนวนหาจำนวนปีที่เป็น leap year กับที่ไม่ใช่ leap year
		yearDiff := ((targetYear - 1) - baseYear)
		// คำนวณหาจำนวน No leap year
		countNotLeapYear := yearDiff - countLeapYear
		// หาจำนวนวันที่ผ่านมาทั้งหมดตั้งแต่ปี base จนถึง ปีก่อนปีเป้าหมาย โดยนำจำนวน leap year * 366 และ No leap year * 365 แล้วบวกกัน
		sumDayPastYear = (countNotLeapYear * 365) + (countLeapYear * 366)
	}
	// นับจำนวณวันที่ผ่านมาทั้งหมดของปีเป้าหมาย
	sumDayTargetYear, err := sumDayOfTargetYear(targetYear, targetMonth, targetDay)
	if err != nil {
		return "", err
	}
	// นำจำนวนวันที่ผ่านมาทั้งหมดตั้งแต่ปี base จนถึง ปีก่อนปีเป้าหมาย บวกกับ จำนวณวันที่ผ่านมาทั้งหมดของปีเป้าหมาย เพื่อหาจำนวนวันทั้งหมด
	sumDay := sumDayTargetYear + sumDayPastYear
	var result string
	if sumDay < 7 {
		// ถ้า sumDay < 7 จะทำให้ mod ค่าไม่ถูกต้อง เลยต้องเอา sumDay - 1 ใช้เป็น index ของ weekday
		result = weekDay[int(sumDay)-1]
	} else {
		// นำจำนวณวันทั้งหมด mod ด้วย 7 จะได้ index ของ weekday
		result = weekDay[int(math.Mod(sumDay, 7))]
	}
	return result, nil
}

func calCountOfLeapYear(baseYear float64, targetYear float64) (float64, error) {

	// เช็คว่า baseYear กับ targetYear ต้องมากกว่าหรือเท่ากับ 0
	if baseYear < 0 || targetYear < 0 {
		return 0, errors.New("baseYear and targetYear must eqaul 0 or more than 0")
	}
	// เช็คว่า baseYear กับ targetYear ต้องมากกว่าหรือเท่ากับ 0
	if baseYear >= targetYear {
		return 0, errors.New("baseYear must less than targetYear")
	}

	countLeapYear := float64(0)
	// หาปีแรกที่มีโอกาสเป็น leap year ของช่วงเวลาระหว่าง baseYear กับ targetYear
	for math.Mod(baseYear, 4) != 0 {
		baseYear++
	}
	// นำปีแรกที่มีโอกาสเป็น leap year + 4 ไปเรื่อยๆเพื่อหาปีที่มีโอกาสเป็น leap year ถัดไป(+ 4 เพราะ leap year มีโอกาสเกิดขึ้นทุก 4 ปี)
	for baseYear <= targetYear {
		// เช็ดว่าจะเป็น leap year หรือไม่
		if isLeapYear(baseYear) {
			//นับจำนวนปีที่เป็น leap year
			countLeapYear++
		}
		baseYear = baseYear + 4
	}
	return countLeapYear, nil
}

func sumDayOfTargetYear(targetYear float64, targetMonth float64, targetDay float64) (float64, error) {
	// เช็ควันที่ว่าตรงกับความจริงไหม
	if isDateOutOfLenght(targetYear, targetMonth, targetDay) {
		return 0, errDateOutOfLenght
	}

	sumDay := float64(0)
	// for loop หาจำนวนวันของเดือนที่ผ่านมา
	for monthIndex := 0; monthIndex < int(targetMonth)-1; monthIndex++ {
		// ถ้า monthIndex == 1 (เดือน 2)
		if monthIndex == 1 {
			// เช็ดว่าจะเป็น leap year หรือไม่
			if isLeapYear(targetYear) {
				//ถ้าเป็น leap year เดือน 2 จะมี 29 วัน
				sumDay += 29
			} else {
				//ถ้าไม่เป็น leap year เดือน 2 จะมี 28 วัน
				sumDay += 28
			}
		} else {
			// ดึงของมูลวันทั้งหมดของเดือนนั้นมารวมกัน
			sumDay += float64(dayOfMonth[monthIndex])
		}
	}
	// รวมจำนวนวันของเดือนที่ผ่านมากับจำนวนว่าที่ผ่านมาของเดือนเป้าหมาย
	return sumDay + targetDay, nil
}

func isLeapYear(year float64) bool {
	if year < 0 {
		return false
	}
	// ถ้า 4 หารปีนั้นๆลงตัวจะมีโอกาสเป็น leap year
	if math.Mod(year, float64(4)) == 0 {
		// ถ้า 100 หารปีนั้นๆไม่ลงตัวจะเป็น leap year
		if math.Mod(year, float64(100)) != 0 {
			return true
			// ถ้า 400 หารปีนั้นๆลงตัวจะเป็น leap year
		} else if math.Mod(year, float64(400)) == 0 {
			return true
		}
	}
	return false
}

func isDateOutOfLenght(targetYear float64, targetMonth float64, targetDay float64) bool {
	//ถ้าเดือน > 12 หรือ เดือน < 1 วันที่จะไม่เป็นไปตามความเป็นจริง
	if targetMonth > 12 || targetMonth < 1 {
		return true
	}
	var maxDayOfMonth int
	// หาจำนวนวันสูงสุดของเดือนนั้นๆ
	if targetMonth == 2 {
		if isLeapYear(targetYear) {
			maxDayOfMonth = 29
		} else {
			maxDayOfMonth = 28
		}
	} else {
		maxDayOfMonth = dayOfMonth[int(targetMonth)-1]
	}
	//ถ้าวันที่มากกว่าจำนวนวันสูงสุดของเดือนนั้นๆ วันที่จะไม่เป็นไปตามความเป็นจริง
	return int(targetDay) > maxDayOfMonth
}
