// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"errors"
	"strconv"
	"time"

	"github.com/guanguans/id-validator/data"
)

// IdInfo 身份证信息
type IdInfo struct {
	AddressCode   int
	Abandoned     int
	Address       string
	AddressTree   []string
	Birthday      time.Time
	Constellation string
	ChineseZodiac string
	Sex           int
	Length        int
	CheckBit      string
}

// IsValid 验证身份证号合法性
func IsValid(id string, strict bool) bool {
	code, err := generateCode(id)
	if err != nil {
		return false
	}

	// 检查顺序码、生日码、地址码
	if !checkOrderCode(code["order"]) || !checkBirthdayCode(code["birthdayCode"]) || !checkAddressCode(code["addressCode"], code["birthdayCode"], strict) {
		return false
	}

	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 校验码
	return code["checkBit"] == generatorCheckBit(code["body"])
}

// GetInfo 获取身份证信息
func GetInfo(id string, strict bool) (IdInfo, error) {
	// 验证有效性
	if !IsValid(id, strict) {
		return IdInfo{}, errors.New("not Valid ID card number")
	}

	code, _ := generateCode(id)
	addressCode, _ := strconv.Atoi(code["addressCode"])

	// 地址信息
	addressInfo := getAddressInfo(code["addressCode"], code["birthdayCode"], strict)
	var addressTree []string //nolint:prealloc
	for _, val := range addressInfo {
		addressTree = append(addressTree, val)
	}

	// 是否废弃
	var abandoned int
	if data.AddressCode[addressCode] == "" {
		abandoned = 1
	}

	// 生日
	cst, _ := time.LoadLocation("Asia/Shanghai")
	birthday, _ := time.ParseInLocation("20060102", code["birthdayCode"], cst)

	// 性别
	sex := 1
	sexCode, _ := strconv.Atoi(code["order"])
	if (sexCode % 2) == 0 {
		sex = 0
	}

	// 长度
	length, _ := strconv.Atoi(code["type"])

	return IdInfo{
		AddressCode:   addressCode,
		Abandoned:     abandoned,
		Address:       addressInfo["province"] + addressInfo["city"] + addressInfo["district"],
		AddressTree:   addressTree,
		Birthday:      birthday,
		Constellation: getConstellation(code["birthdayCode"]),
		ChineseZodiac: getChineseZodiac(code["birthdayCode"]),
		Sex:           sex,
		Length:        length,
		CheckBit:      code["checkBit"],
	}, nil
}

// FakeId 生成假身份证号码
func FakeId() string {
	return FakeRequireId(true, "", "", 0)
}

// FakeRequireId 按要求生成假身份证号码
// isEighteen 是否生成18位号码
// address    省市县三级地区官方全称：如`北京市`、`台湾省`、`香港特别行政区`、`深圳市`、`黄浦区`
// birthday   出生日期：如 `2000`、`198801`、`19990101`
// sex        性别：1为男性，0为女性
func FakeRequireId(isEighteen bool, address string, birthday string, sex int) string {

	// 生成地址码
	var addressCode string
	if address == "" {
		for i, s := range data.AddressCode {
			addressCode = strconv.Itoa(i)
			address = s
			break
		}
	} else {
		addressCode = generatorAddressCode(address)
	}

	// 出生日期码
	birthdayCode := generatorBirthdayCode(addressCode, address, birthday)

	// 生成顺序码
	orderCode := generatorOrderCode(sex)

	if !isEighteen {
		return addressCode + substr(birthdayCode, 2, 8) + orderCode
	}

	body := addressCode + birthdayCode + orderCode

	return body + generatorCheckBit(body)
}

// UpgradeId 15位升级18位号码
func UpgradeId(id string) (string, error) {
	if !IsValid(id, true) {
		return "", errors.New("not Valid ID card number")
	}

	code, _ := generateShortCode(id)

	body := code["addressCode"] + code["birthdayCode"] + code["order"]

	return body + generatorCheckBit(body), nil
}
