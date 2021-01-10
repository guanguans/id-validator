package id_validator

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"id-validator/data"
)

// 生成 Bit 码
func GeneratorCheckBit(body string) string {
	// 位置加权
	var posWeight [19]float64
	for i := 2; i < 19; i++ {
		weight := int(math.Pow(2, float64(i-1))) % 11
		posWeight[i] = float64(weight)
	}

	// 累身份证号 body 部分与位置加权的积
	bodySum := 0
	bodyArray := strings.Split(body, "")
	count := len(bodyArray)
	for i := 0; i < count; i++ {
		bodySubStr, _ := strconv.Atoi(bodyArray[i])
		bodySum += bodySubStr * int(posWeight[18-i])
	}

	// 生成校验码
	checkBit := (12 - (bodySum % 11)) % 11
	if checkBit == 10 {
		return "X"
	}
	return strconv.Itoa(checkBit)
}

// 生成地址码
func GeneratorAddressCode(address string) string {
	var addressCodeInt int
	for code, addressStr := range data.AddressCode {
		if address == addressStr {
			addressCodeInt = code
			break
		}
	}
	addressCode := strconv.Itoa(addressCodeInt)

	classification := AddressCodeClassification(strconv.Itoa(addressCodeInt))

	switch classification {
	case "country":
		// addressCode = GetRandAddressCode("\\d{4}(?!00)[0-9]{2}$")
		addressCode = GetRandAddressCode("\\d{4}(?)[0-9]{2}$")
	case "province":
		provinceCode := Substr(addressCode, 0, 2)
		// pattern := "^" + provinceCode + "\\d{2}(?!00)[0-9]{2}$"
		pattern := "^" + provinceCode + "\\d{2}(?)[0-9]{2}$"
		addressCode = GetRandAddressCode(pattern)
	case "city":
		cityCode := Substr(addressCode, 0, 4)
		// pattern := "^" + cityCode + "(?!00)[0-9]{2}$"
		pattern := "^" + cityCode + "(?)[0-9]{2}$"
		addressCode = GetRandAddressCode(pattern)
	}

	return addressCode
}

// 地址码分类
func AddressCodeClassification(addressCode string) string {
	// 全国
	if addressCode == "" {
		return "country"
	}

	// 港澳台
	if Substr(addressCode, 0, 1) == "8" {
		return "special"
	}

	// 省级
	if Substr(addressCode, 2, 6) == "0000" {
		return "province"
	}

	// 市级
	if Substr(addressCode, 4, 6) == "00" {
		return "city"
	}

	// 县级
	return "district"
}

// 获取随机地址码
func GetRandAddressCode(pattern string) string {
	mustCompile := regexp.MustCompile(pattern)

	var keys []string
	for key, _ := range data.AddressCode {
		keyStr := strconv.Itoa(key)
		if mustCompile.MatchString(keyStr) && Substr(keyStr, 4, 6) != "00" {
			keys = append(keys, keyStr)
		}
	}

	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	randKey := rand.Intn(len(keys))

	return keys[randKey]
}

// 生成出生日期码
func GeneratorBirthdayCode(birthday string) string {
	year, _ := strconv.Atoi(DatePad(Substr(birthday, 0, 4), "year"))
	month, _ := strconv.Atoi(DatePad(Substr(birthday, 4, 6), "month"))
	day, _ := strconv.Atoi(DatePad(Substr(birthday, 6, 8), "day"))

	nowYear := time.Now().Year()
	rand.Seed(time.Now().Unix())
	if year < 1800 || year > nowYear {
		randYear := rand.Intn(nowYear-1950) + 1950
		year, _ = strconv.Atoi(DatePad(strconv.Itoa(randYear), "year"))
	}
	if month < 1 || month > 12 {
		randMonth := rand.Intn(12-1) + 1
		month, _ = strconv.Atoi(DatePad(strconv.Itoa(randMonth), "month"))
	}
	if day < 1 || day > 31 {
		randDay := rand.Intn(28-1) + 1
		day, _ = strconv.Atoi(DatePad(strconv.Itoa(randDay), "day"))
	}

	birthdayStr := strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(day)
	_, error := time.Parse("20060102", birthdayStr)
	if error != nil {
		randYear := rand.Intn(nowYear-1950) + 1950
		year, _ = strconv.Atoi(DatePad(strconv.Itoa(randYear), "year"))

		randMonth := rand.Intn(12-1) + 1
		month, _ = strconv.Atoi(DatePad(strconv.Itoa(randMonth), "month"))

		randDay := rand.Intn(28-1) + 1
		day, _ = strconv.Atoi(DatePad(strconv.Itoa(randDay), "day"))
	}

	return strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(day)
}

// 生成顺序码
func GeneratorOrderCode(sex int) string {
	rand.Seed(time.Now().Unix())
	orderCode := rand.Intn(999-111) + 111
	if sex != orderCode%2 {
		orderCode -= 1
	}

	return strconv.Itoa(orderCode)
}

// 日期补全
func DatePad(date string, category string) string {
	padLength := 2
	if category == "year" {
		padLength = 4
	}

	return fmt.Sprintf("%s%032s", date, "")[:padLength]
}
