package id_validator

import (
	"fmt"
	"strconv"

	"id-validator/data"
)

// 验证身份证号合法性
func IsValid(id string) bool {
	if !CheckIdArgument(id) {
		return false
	}

	code := GenerateType(id)
	if !CheckAddressCode(code["addressCode"], code["birthdayCode"]) || !CheckBirthdayCode(code["birthdayCode"]) || !CheckOrderCode(code["order"]) {
		return false
	}

	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 验证：校验码
	checkBit := GeneratorCheckBit(code["body"])

	return code["checkBit"] == checkBit
}

// 获取身份证信息
func GetInfo(id string) map[string]string {
	// 验证有效性
	if !IsValid(id) {
		return map[string]string{}
	}

	code := GenerateType(id)

	addressInfo := GetAddressInfo(code["addressCode"], code["birthdayCode"])
	fmt.Println(addressInfo)
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
