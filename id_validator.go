package idvalidator

import (
	"errors"
	"strconv"
	"time"

	"id-validator/data"
)

// 身份证信息
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
	CheckBit      int
}

// 验证身份证号合法性
func IsValid(id string) bool {
	code, err := GenerateCode(id)
	if err != nil {
		return false
	}

	// 检查顺序码、生日码、地址码
	if !CheckOrderCode(code["order"]) || !CheckBirthdayCode(code["birthdayCode"]) || !CheckAddressCode(code["addressCode"], code["birthdayCode"]) {
		return false
	}

	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 校验码
	return code["checkBit"] == GeneratorCheckBit(code["body"])
}

// 获取身份证信息
func GetInfo(id string) (IdInfo, error) {
	// 验证有效性
	if !IsValid(id) {
		return IdInfo{}, errors.New("Not Valid ID card number.")
	}

	code, _ := GenerateCode(id)
	addressCode, _ := strconv.Atoi(code["addressCode"])

	// 地址信息
	addressInfo := GetAddressInfo(code["addressCode"], code["birthdayCode"])
	var addressTree []string
	for _, val := range addressInfo {
		addressTree = append(addressTree, val)
	}

	// 是否废弃
	var abandoned int
	if data.AddressCode[addressCode] == "" {
		abandoned = 1
	}

	// 生日
	birthday, _ := time.Parse("20060102", code["birthdayCode"])

	// 性别
	sex := 1
	sexCode, _ := strconv.Atoi(code["order"])
	if (sexCode % 2) == 0 {
		sex = 0
	}

	// 长度
	length, _ := strconv.Atoi(code["type"])

	// Bit码
	checkBit, _ := strconv.Atoi(code["checkBit"])

	return IdInfo{
		AddressCode:   addressCode,
		Abandoned:     abandoned,
		Address:       addressInfo["province"] + addressInfo["city"] + addressInfo["district"],
		AddressTree:   addressTree,
		Birthday:      birthday,
		Constellation: GetConstellation(code["birthdayCode"]),
		ChineseZodiac: GetChineseZodiac(code["birthdayCode"]),
		Sex:           sex,
		Length:        length,
		CheckBit:      checkBit,
	}, nil
}

// 生成假身份证号码
func Fake(isEighteen bool, address string, birthday string, sex int) string {
	// 生成地址码
	addressCode := GeneratorAddressCode(address)

	// 出生日期码
	birthdayCode := GeneratorBirthdayCode(birthday)

	// 生成顺序码
	orderCode := GeneratorOrderCode(sex)

	if !isEighteen {
		return addressCode + Substr(birthdayCode, 2, 6) + orderCode
	}

	body := addressCode + birthdayCode + orderCode

	return body + GeneratorCheckBit(body)
}

// 15位升级18位号码
func Upgrade(id string) (string, error) {
	if !IsValid(id) {
		return "", errors.New("Not Valid ID card number.")
	}

	code, _ := GenerateShortCode(id)

	body := code["addressCode"] + code["birthdayCode"] + code["order"]

	return body + GeneratorCheckBit(body), nil
}
