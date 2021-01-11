package idvalidator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 检查ID参数
func CheckIDArgument(id string) bool {
	_, err := GenerateCode(id)

	return err == nil
}

// 生成数据
func GenerateCode(id string) (map[string]string, error) {
	length := len(id)
	if length == 15 {
		return GenerateShortCode(id)
	}

	if length == 18 {
		return GenerateLongCode(id)
	}

	return map[string]string{}, errors.New("Invalid ID card number length.")
}

// 生成短数据
func GenerateShortCode(id string) (map[string]string, error) {
	if len(id) != 15 {
		return map[string]string{}, errors.New("Invalid ID card number length.")
	}

	mustCompile := regexp.MustCompile("(.{6})(.{6})(.{3})")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))

	return map[string]string{
		"body":         subMatch[0],
		"addressCode":  subMatch[1],
		"birthdayCode": "19" + subMatch[2],
		"order":        subMatch[3],
		"checkBit":     "",
		"type":         "15",
	}, nil
}

// 生成长数据
func GenerateLongCode(id string) (map[string]string, error) {
	if len(id) != 18 {
		return map[string]string{}, errors.New("Invalid ID card number length.")
	}
	mustCompile := regexp.MustCompile("((.{6})(.{8})(.{3}))(.)")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))

	return map[string]string{
		"body":         subMatch[1],
		"addressCode":  subMatch[2],
		"birthdayCode": subMatch[3],
		"order":        subMatch[4],
		"checkBit":     subMatch[5],
		"type":         "18",
	}, nil
}

// 检查地址码
func CheckAddressCode(addressCode string, birthdayCode string) bool {
	return GetAddressInfo(addressCode, birthdayCode)["province"] != ""
}

// 检查出生日期码
func CheckBirthdayCode(birthdayCode string) bool {
	year, _ := strconv.Atoi(Substr(birthdayCode, 0, 4))
	if year < 1800 {
		return false
	}

	nowYear := time.Now().Year()
	if year > nowYear {
		return false
	}

	_, err := time.Parse("20060102", birthdayCode)

	return err == nil
}

// 检查顺序码
func CheckOrderCode(orderCode string) bool {
	return len(orderCode) == 3
}
