package idvalidator

import (
	"id-validator/data"
	"strconv"
	"strings"
)

// 检查地址码
func GetAddressInfo(addressCode string, birthdayCode string) map[string]string {
	addressInfo := map[string]string{
		"province": "",
		"city":     "",
		"district": "",
	}

	// 省级信息
	addressInfo["province"] = GetAddress(Substr(addressCode, 0, 2)+"0000", birthdayCode)

	// 用于判断是否是港澳台居民居住证（8字开头）
	firstCharacter := Substr(addressCode, 0, 1)
	// 港澳台居民居住证无市级、县级信息
	if firstCharacter == "8" {
		return addressInfo
	}

	// 市级信息
	addressInfo["city"] = GetAddress(Substr(addressCode, 0, 4)+"00", birthdayCode)

	// 县级信息
	addressInfo["district"] = GetAddress(addressCode, birthdayCode)

	return addressInfo
}

// 获取省市区地址码
func GetAddress(addressCode string, birthdayCode string) string {
	address := ""
	addressCodeInt, _ := strconv.Atoi(addressCode)
	year, _ := strconv.Atoi(Substr(birthdayCode, 0, 4))
	for key, val := range data.AddressCodeTimeline[addressCodeInt] {
		// if len(val) == 0 {
		// 	continue
		// }
		startYear, _ := strconv.Atoi(val["start_year"])
		if (key == 0 && year < startYear) || year >= startYear {
			address = val["address"]
		}
	}

	return address
}

// 获取星座信息
func GetConstellation(birthdayCode string) string {
	monthStr := Substr(birthdayCode, 4, 6)
	dayStr := Substr(birthdayCode, 6, 8)
	month, _ := strconv.Atoi(monthStr)
	day, _ := strconv.Atoi(dayStr)
	startDate := data.Constellation[month]["start_date"]
	startDay, _ := strconv.Atoi(strings.Split(startDate, "-")[1])
	if day >= startDay {
		return data.Constellation[month]["name"]
	}

	tmpMonth := month - 1
	if month == 1 {
		tmpMonth = 12
	}

	return data.Constellation[tmpMonth]["name"]
}

// 获取生肖信息
func GetChineseZodiac(birthdayCode string) string {
	// 子鼠
	start := 1900
	end, _ := strconv.Atoi(Substr(birthdayCode, 0, 4))
	key := (end - start) % 12

	return data.ChineseZodiac[key]
}

// Substr 截取字符串
func Substr(source string, start int, end int) string {
	r := []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
