// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/guanguans/id-validator/data"
	"github.com/spf13/cast"
)

// 生成Bit码
func generatorCheckBit(body string) string {
	// 位置加权
	var posWeight [19]float64
	for i := 2; i < 19; i++ {
		weight := int(math.Pow(2, float64(i-1))) % 11
		posWeight[i] = float64(weight)
	}

	// 累身份证号body部分与位置加权的积
	var bodySum int
	bodyArray := strings.Split(body, "")
	count := len(bodyArray)
	for i := 0; i < count; i++ {
		bodySum += cast.ToInt(bodyArray[i]) * int(posWeight[18-i])
	}

	// 生成校验码
	checkBit := (12 - (bodySum % 11)) % 11
	if checkBit == 10 {
		return "x"
	}
	return cast.ToString(checkBit)
}

// 生成地址码
func generatorAddressCode(address string) string {
	addressCode := ""
	for code, addressStr := range data.AddressCode() {
		if address == addressStr {
			addressCode = cast.ToString(code)
			break
		}
	}

	classification := addressCodeClassification(addressCode)
	switch classification {
	case "country":
		addressCode = getRandAddressCode("\\d{4}(?)[0-9]{2}$")
	case "province":
		pattern := "^" + substr(addressCode, 0, 2) + "\\d{2}(?)[0-9]{2}$"
		addressCode = getRandAddressCode(pattern)
	case "city":
		pattern := "^" + substr(addressCode, 0, 4) + "(?)[0-9]{2}$"
		addressCode = getRandAddressCode(pattern)
	}

	return addressCode
}

// 地址码分类
func addressCodeClassification(addressCode string) string {
	// 全国
	if addressCode == "" {
		return "country"
	}

	// 港澳台
	if substr(addressCode, 0, 1) == "8" {
		return "special"
	}

	// 省级
	if substr(addressCode, 2, 6) == "0000" {
		return "province"
	}

	// 市级
	if substr(addressCode, 4, 6) == "00" {
		return "city"
	}

	// 县级
	return "district"
}

// 获取随机地址码
func getRandAddressCode(pattern string) string {
	mustCompile := regexp.MustCompile(pattern)
	var keys []string
	for key := range data.AddressCode() {
		keyStr := cast.ToString(key)
		if mustCompile.MatchString(keyStr) && substr(keyStr, 4, 6) != "00" {
			keys = append(keys, keyStr)
		}
	}

	rand.Seed(time.Now().Unix())

	return keys[rand.Intn(len(keys))]
}

// 生成出生日期码
func generatorBirthdayCode(addressCode string, address string, birthday string) string {
	startYear := "0001"
	endYear := "9999"
	year := datePipeHandle(datePad(substr(birthday, 0, 4), "year"), "year")
	month := datePipeHandle(datePad(substr(birthday, 4, 6), "month"), "month")
	day := datePipeHandle(datePad(substr(birthday, 6, 8), "day"), "day")

	if timeline := data.GetAddressCodeTimeline(cast.ToUint32(addressCode)); len(timeline) > 0 {
		for _, val := range timeline {
			if val["address"] == address {
				if val["start_year"] != "" {
					startYear = val["start_year"]
				}
				if val["end_year"] != "" {
					endYear = val["end_year"]
				}
			}
		}
	}

	yearInt := cast.ToInt(year)
	if yearInt < cast.ToInt(startYear) {
		year = startYear
	}
	if yearInt > cast.ToInt(endYear) {
		year = endYear
	}

	birthday = year + month + day
	_, err := time.Parse("20060102", birthday)
	// example: 195578
	if err != nil {
		year = datePad(year, "year")
		month = datePad(month, "month")
		day = datePad(day, "day")
	}

	return year + month + day
}

// 日期处理
func datePipeHandle(date string, category string) string {
	dateInt := cast.ToInt(date)

	switch category {
	case "year":
		nowYear := time.Now().Year()
		rand.Seed(time.Now().Unix())
		if dateInt < 1800 || dateInt > nowYear {
			randDate := rand.Intn(nowYear-1950) + 1950
			date = cast.ToString(randDate)
		}
	case "month":
		if dateInt < 1 || dateInt > 12 {
			randDate := rand.Intn(12-1) + 1
			date = cast.ToString(randDate)
		}

	case "day":
		if dateInt < 1 || dateInt > 31 {
			randDate := rand.Intn(28-1) + 1
			date = cast.ToString(randDate)
		}
	}

	return date
}

// 生成顺序码
func generatorOrderCode(sex int) string {
	rand.Seed(time.Now().Unix())
	orderCode := rand.Intn(999-111) + 111
	if sex != orderCode%2 {
		orderCode--
	}

	return cast.ToString(orderCode)
}

// 日期补全
func datePad(date string, category string) string {
	padLength := 2
	if category == "year" {
		padLength = 4
	}

	return fmt.Sprintf(fmt.Sprintf("%%0%ds", padLength), date)
}
