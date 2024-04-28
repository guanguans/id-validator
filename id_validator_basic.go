package idvalidator

import (
	"errors"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

// BasicIdInfo 不含地址码的身份证信息
type BasicIdInfo struct {
	Birthday      time.Time
	Constellation string
	ChineseZodiac string
	Sex           int
	Length        int
	CheckBit      string
}

// IsValidBasic 验证身份证号除地址信息之外的合法性
func IsValidBasic(id string) bool {
	code, err := generateCode(id)
	if err != nil {
		return false
	}

	// 检查顺序码、生日码、地址码
	if !checkOrderCode(code["order"]) || !checkBirthdayCode(code["birthdayCode"]) {
		return false
	}

	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 校验码
	return code["checkBit"] == generatorCheckBit(code["body"])
}

// GetBasicInfo 获取除地址信息之外的身份证信息
func GetBasicInfo(id string) (BasicIdInfo, error) {
	// 验证有效性
	if !IsValidBasic(id) {
		return BasicIdInfo{}, errors.New("invalid ID card number")
	}

	code, _ := generateCode(id)

	// 生日
	cst, _ := time.LoadLocation("Asia/Shanghai")
	birthday, _ := time.ParseInLocation("20060102", code["birthdayCode"], cst)

	// 性别
	sex := 1
	order, _ := strconv.Atoi(code["order"])
	if (order % 2) == 0 {
		sex = 0
	}

	// 长度
	length := cast.ToInt(code["type"])

	return BasicIdInfo{
		Birthday:      birthday,
		Constellation: getConstellation(code["birthdayCode"]),
		ChineseZodiac: getChineseZodiac(code["birthdayCode"]),
		Sex:           sex,
		Length:        length,
		CheckBit:      code["checkBit"],
	}, nil
}
