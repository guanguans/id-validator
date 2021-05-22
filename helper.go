package idvalidator

import (
	"strconv"
	"strings"

	"github.com/guanguans/id-validator/data"
)

// 获取地址信息
func getAddressInfo(addressCode string, birthdayCode string) map[string]string {
	addressInfo := map[string]string{
		"province": "",
		"city":     "",
		"district": "",
	}

	// 省级信息
	addressInfo["province"] = getAddress(substr(addressCode, 0, 2)+"0000", birthdayCode)

	// 用于判断是否是港澳台居民居住证（8字开头）
	firstCharacter := substr(addressCode, 0, 1)
	// 港澳台居民居住证无市级、县级信息
	if firstCharacter == "8" {
		return addressInfo
	}

	// 市级信息
	addressInfo["city"] = getAddress(substr(addressCode, 0, 4)+"00", birthdayCode)

	// 县级信息
	addressInfo["district"] = getAddress(addressCode, birthdayCode)

	return addressInfo
}

// 获取省市区地址码
func getAddress(addressCode string, birthdayCode string) string {
	address := ""
	addressCodeInt, _ := strconv.Atoi(addressCode)
	year, _ := strconv.Atoi(substr(birthdayCode, 0, 4))
	startYear := "0001"
	endYear := "9999"
	for _, val := range data.AddressCodeTimeline[addressCodeInt] {
		if val["start_year"] != "" {
			startYear = val["start_year"]
		}
		if val["end_year"] != "" {
			endYear = val["end_year"]
		}
		startYearInt, _ := strconv.Atoi(startYear)
		endYearInt, _ := strconv.Atoi(endYear)
		if year >= startYearInt && year <= endYearInt {
			address = val["address"]
		}
	}

	return address
}

// 获取星座信息
func getConstellation(birthdayCode string) string {
	monthStr := substr(birthdayCode, 4, 6)
	dayStr := substr(birthdayCode, 6, 8)
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
func getChineseZodiac(birthdayCode string) string {
	// 子鼠
	start := 1900
	end, _ := strconv.Atoi(substr(birthdayCode, 0, 4))
	key := (end - start) % 12

	return data.ChineseZodiac[key]
}

// substr 截取字符串
func substr(source string, start int, end int) string {
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
