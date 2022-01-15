// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"strconv"
	"strings"

	"github.com/guanguans/id-validator/data"
	"github.com/spf13/cast"
)

// 获取地址信息
func getAddressInfo(addressCode string, birthdayCode string, strict bool) map[string]string {
	addressInfo := map[string]string{
		"province": "",
		"city":     "",
		"district": "",
	}

	// 省级信息
	addressInfo["province"] = getAddress(substr(addressCode, 0, 2)+"0000", birthdayCode, strict)

	// 用于判断是否是港澳台居民居住证（8字开头）
	firstCharacter := substr(addressCode, 0, 1)
	// 港澳台居民居住证无市级、县级信息
	if firstCharacter == "8" {
		return addressInfo
	}

	// 市级信息
	addressInfo["city"] = getAddress(substr(addressCode, 0, 4)+"00", birthdayCode, strict)

	// 县级信息
	addressInfo["district"] = getAddress(addressCode, birthdayCode, strict)

	return addressInfo
}

// 获取省市区地址码
func getAddress(addressCode string, birthdayCode string, strict bool) string {
	address := ""
	addressCodeInt := cast.ToInt(addressCode)
	if _, ok := data.AddressCodeTimeline[addressCodeInt]; !ok {
		// 修复 \d\d\d\d01、\d\d\d\d02、\d\d\d\d11 和 \d\d\d\d20 的历史遗留问题
		// 以上四种地址码，现实身份证真实存在，但民政部历年公布的官方地址码中可能没有查询到
		// 如：440401 450111 等
		// 所以这里需要特殊处理
		// 1980年、1982年版本中，未有制定省辖市市辖区的代码，所有带县的省辖市给予“××××20”的“市区”代码。
		// 1984年版本开始对地级市（前称省辖市）市辖区制定代码，其中“××××01”表示市辖区的汇总码，同时撤销“××××20”的“市区”代码（追溯至1983年）。
		// 1984年版本的市辖区代码分为城区和郊区两类，城区由“××××02”开始排起，郊区由“××××11”开始排起，后来版本已不再采用此方式，已制定的代码继续沿用。
		suffixes := substr("123456", 4, 6)
		switch suffixes {
		case "20":
			address = "市区"
		case "01":
			address = "市辖区"
		case "02":
			address = "城区"
		case "11":
			address = "郊区"
		}

		return address
	}

	timeline := data.AddressCodeTimeline[addressCodeInt]
	year := cast.ToInt(substr(birthdayCode, 0, 4))
	startYear := "0001"
	endYear := "9999"
	for _, val := range timeline {
		if val["start_year"] != "" {
			startYear = val["start_year"]
		}
		if val["end_year"] != "" {
			endYear = val["end_year"]
		}
		if year >= cast.ToInt(startYear) && year <= cast.ToInt(endYear) {
			address = val["address"]
		}
	}

	if address == "" && !strict {
		// 由于较晚申请户口或身份证等原因，导致会出现地址码正式启用于2000年，但实际1999年出生的新生儿，由于晚了一年报户口，导致身份证上的出生年份早于地址码正式启用的年份
		// 由于某些地区的地址码已经废弃，但是实际上在之后的几年依然在使用
		// 这里就不做时间判断了
		address = timeline[0]["address"]
	}

	return address
}

// 获取星座信息
func getConstellation(birthdayCode string) string {
	month, _ := strconv.Atoi(substr(birthdayCode, 4, 6))
	day, _ := strconv.Atoi(substr(birthdayCode, 6, 8))
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
	end := cast.ToInt(substr(birthdayCode, 0, 4))
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
