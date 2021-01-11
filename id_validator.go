package idvalidator

import (
	"errors"
	"id-validator/data"
	"strconv"
)

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
func GetInfo(id string) map[string]string {
	// 验证有效性
	if !IsValid(id) {
		return map[string]string{}
	}

	code, _ := GenerateCode(id)

	addressInfo := GetAddressInfo(code["addressCode"], code["birthdayCode"])
	// fmt.Println(addressInfo)
	address, _ := strconv.Atoi(code["addressCode"])
	abandoned := "0"
	if data.AddressCode[address] == "" {
		abandoned = "1"
	}
	// birthday, _ := time.Parse("20060102", code["birthdayCode"])

	sex := "1"
	sexCode, _ := strconv.Atoi(code["order"])
	if (sexCode % 2) == 0 {
		sex = "0"
	}
	info := map[string]string{
		"addressCode": code["addressCode"],
		"abandoned":   abandoned,
		"address":     addressInfo["province"] + addressInfo["city"] + addressInfo["district"],
		// "addressTree": addressInfo,
		// "birthdayCode": birthday,
		"constellation": GetConstellation(code["birthdayCode"]),
		"chineseZodiac": GetChineseZodiac(code["birthdayCode"]),
		"sex":           sex,
		"length":        code["type"],
		"checkBit":      code["checkBit"],
	}

	return info
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
