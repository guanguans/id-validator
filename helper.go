package id_validator

import (
	"strconv"
	"strings"

	"id-validator/data"
)

// 检查地址码
func GetAddressInfo(addressCode string, birthdayCode string) map[string]string {
	addressInfo := map[string]string{
		"province": "",
		"city":     "",
		"district": "",
	}

	// 省级信息
	provinceAddressCode := Substr(addressCode, 0, 2) + "0000"
	addressInfo["province"] = GetAddress(provinceAddressCode, birthdayCode)

	// 用于判断是否是港澳台居民居住证（8字开头）
	firstCharacter := Substr(addressCode, 0, 1)
	// 港澳台居民居住证无市级、县级信息
	if firstCharacter == "8" {
		return addressInfo
	}

	// 市级信息
	cityAddressCode := Substr(addressCode, 0, 4) + "00"
	addressInfo["city"] = GetAddress(cityAddressCode, birthdayCode)

	// 县级信息
	addressInfo["district"] = GetAddress(addressCode, birthdayCode)

	return addressInfo
}

func GetAddress(addressCode string, birthdayCode string) string {
	var address string

	addressCodeStr, _ := strconv.Atoi(addressCode)
	addressCodeTimeline := data.AddressCodeTimeline[addressCodeStr]

	year := Substr(birthdayCode, 0, 4)
	yearStr, _ := strconv.Atoi(year)

	for key, val := range addressCodeTimeline {
		if len(val) == 0 {
			continue
		}

		startYear, _ := strconv.Atoi(val["start_year"])
		if (key == 0 && yearStr < startYear) || yearStr >= startYear {
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
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
