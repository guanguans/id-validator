package idvalidator

import (
	"id-validator/data"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 生成Bit码
func GeneratorCheckBit(body string) string {
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
		bodySub, _ := strconv.Atoi(bodyArray[i])
		bodySum += bodySub * int(posWeight[18-i])
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
	addressCode := ""
	for code, addressStr := range data.AddressCode {
		if address == addressStr {
			addressCode = strconv.Itoa(code)
			break
		}
	}

	classification := AddressCodeClassification(addressCode)
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
	for key := range data.AddressCode {
		keyStr := strconv.Itoa(key)
		if mustCompile.MatchString(keyStr) && Substr(keyStr, 4, 6) != "00" {
			keys = append(keys, keyStr)
		}
	}

	rand.Seed(time.Now().Unix())

	return keys[rand.Intn(len(keys))]
}

// 生成出生日期码
func GeneratorBirthdayCode(birthday string) string {
	year := DatePipelineHandle(DatePad(Substr(birthday, 0, 4), "year"), "year")
	month := DatePipelineHandle(DatePad(Substr(birthday, 4, 6), "month"), "month")
	day := DatePipelineHandle(DatePad(Substr(birthday, 6, 8), "day"), "day")

	birthday = year + month + day
	_, error := time.Parse("20060102", birthday)
	// example: 195578
	if error != nil {
		year = DatePad(year, "year")
		month = DatePad(month, "month")
		day = DatePad(day, "day")
	}

	return year + month + day
}

// 日期处理
func DatePipelineHandle(date string, category string) string {
	dateInt, _ := strconv.Atoi(date)

	switch category {
	case "year":
		nowYear := time.Now().Year()
		rand.Seed(time.Now().Unix())
		if dateInt < 1800 || dateInt > nowYear {
			randDate := rand.Intn(nowYear-1950) + 1950
			date = strconv.Itoa(randDate)
		}
	case "month":
		if dateInt < 1 || dateInt > 12 {
			randDate := rand.Intn(12-1) + 1
			date = strconv.Itoa(randDate)
		}

	case "day":
		if dateInt < 1 || dateInt > 31 {
			randDate := rand.Intn(28-1) + 1
			date = strconv.Itoa(randDate)
		}
	}

	return date
}

// 生成顺序码
func GeneratorOrderCode(sex int) string {
	rand.Seed(time.Now().Unix())
	orderCode := rand.Intn(999-111) + 111
	if sex != orderCode%2 {
		orderCode--
	}

	return strconv.Itoa(orderCode)
}

// 日期补全
func DatePad(date string, category string) string {
	padLength := 2
	if category == "year" {
		padLength = 4
	}

	for i := 0; i < padLength; i++ {
		length := len([]rune(date))
		if length < padLength {
			// date = fmt.Sprintf("%s%s", "0", date)
			date = "0" + date
		}
	}

	return date
}
