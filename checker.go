// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// 检查ID参数
func checkIDArgument(id string) bool {
	_, err := generateCode(id)

	return err == nil
}

// 生成数据
func generateCode(id string) (map[string]string, error) {
	length := len(id)
	if length == 15 {
		return generateShortCode(id)
	}

	if length == 18 {
		return generateLongCode(id)
	}

	return map[string]string{}, errors.New("invalid ID card number length")
}

// 生成短数据
func generateShortCode(id string) (map[string]string, error) {
	if len(id) != 15 {
		return map[string]string{}, errors.New("invalid ID card number length")
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
func generateLongCode(id string) (map[string]string, error) {
	if len(id) != 18 {
		return map[string]string{}, errors.New("invalid ID card number length")
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
func checkAddressCode(addressCode string, birthdayCode string, strict bool) bool {
	return getAddressInfo(addressCode, birthdayCode, strict)["province"] != ""
}

// 检查出生日期码
func checkBirthdayCode(birthdayCode string) bool {
	year := cast.ToInt(substr(birthdayCode, 0, 4))
	if year < 1800 {
		return false
	}

	if year > time.Now().Year() {
		return false
	}

	_, err := time.Parse("20060102", birthdayCode)

	return err == nil
}

// 检查顺序码
func checkOrderCode(orderCode string) bool {
	return len(orderCode) == 3
}
