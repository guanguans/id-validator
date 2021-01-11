package id_validator

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 检查ID参数
func CheckIdArgument(id string) bool {
	return len(GenerateType(id)) != 0
}

// 生成数据
func GenerateType(id string) map[string]string {
	lowerId := strings.ToLower(id)

	if len(lowerId) == 15 {
		return GenerateShortType(id)
	}

	if len(lowerId) == 18 {
		return GenerateLongType(id)
	}

	return map[string]string{}
}

// 生成短数据
func GenerateShortType(id string) map[string]string {
	mustCompile := regexp.MustCompile("(.{6})(.{6})(.{3})")
	subMatch := mustCompile.FindStringSubmatch(id)

	return map[string]string{
		"body":         subMatch[0],
		"addressCode":  subMatch[1],
		"birthdayCode": "19" + subMatch[2],
		"order":        subMatch[3],
		"checkBit":     "",
		"type":         "15",
	}
}

// 生成长数据
func GenerateLongType(id string) map[string]string {
	mustCompile := regexp.MustCompile("((.{6})(.{8})(.{3}))(.)")
	subMatch := mustCompile.FindStringSubmatch(id)

	return map[string]string{
		"body":         subMatch[1],
		"addressCode":  subMatch[2],
		"birthdayCode": subMatch[3],
		"order":        subMatch[4],
		"checkBit":     subMatch[5],
		"type":         "18",
	}
}

// 检查地址码
func CheckAddressCode(addressCode string, birthdayCode string) bool {
	addressInfo := GetAddressInfo(addressCode, birthdayCode)

	return addressInfo["province"] != ""
}

// 检查出生日期码
func CheckBirthdayCode(birthdayCode string) bool {
	year, _ := strconv.Atoi(Substr(birthdayCode, 0, 4))
	if year < 1800 {
		return false
	}

	_, error := time.Parse("20060102", birthdayCode)

	return error == nil
}

// 检查顺序码
func CheckOrderCode(orderCode string) bool {
	return len(orderCode) == 3
}
