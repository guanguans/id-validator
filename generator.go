package idvalidator

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/guanguans/id-validator/data"
)

// 生成Bit码
func generatorCheckBit(body string) string {
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
		return "x"
	}
	return strconv.Itoa(checkBit)
}

// 生成地址码
func generatorAddressCode(address string) string {
	addressCode := ""
	for code, addressStr := range data.AddressCode {
		if address == addressStr {
			addressCode = strconv.Itoa(code)
			break
		}
	}

	classification := addressCodeClassification(addressCode)
	switch classification {
	case "country":
		// addressCode = getRandAddressCode("\\d{4}(?!00)[0-9]{2}$")
		addressCode = getRandAddressCode("\\d{4}(?)[0-9]{2}$")
	case "province":
		provinceCode := substr(addressCode, 0, 2)
		// pattern := "^" + provinceCode + "\\d{2}(?!00)[0-9]{2}$"
		pattern := "^" + provinceCode + "\\d{2}(?)[0-9]{2}$"
		addressCode = getRandAddressCode(pattern)
	case "city":
		cityCode := substr(addressCode, 0, 4)
		// pattern := "^" + cityCode + "(?!00)[0-9]{2}$"
		pattern := "^" + cityCode + "(?)[0-9]{2}$"
		addressCode = getRandAddressCode(pattern)
	}

	return addressCode
}

// 地址码分类
func addressCodeClassification(addressCode string) string {
	// 全国
	if addressCode == "" {
		return "country"
	}

	// 港澳台
	if substr(addressCode, 0, 1) == "8" {
		return "special"
	}

	// 省级
	if substr(addressCode, 2, 6) == "0000" {
		return "province"
	}

	// 市级
	if substr(addressCode, 4, 6) == "00" {
		return "city"
	}

	// 县级
	return "district"
}

// 获取随机地址码
func getRandAddressCode(pattern string) string {
	mustCompile := regexp.MustCompile(pattern)
	var keys []string
	for key := range data.AddressCode {
		keyStr := strconv.Itoa(key)
		if mustCompile.MatchString(keyStr) && substr(keyStr, 4, 6) != "00" {
			keys = append(keys, keyStr)
		}
	}

	rand.Seed(time.Now().Unix())

	return keys[rand.Intn(len(keys))]
}

// 生成出生日期码
func generatorBirthdayCode(addressCode string, address string, birthday string) string {
	startYaer := "0001"
	endYear := "9999"
	year := datePipelineHandle(datePad(substr(birthday, 0, 4), "year"), "year")
	month := datePipelineHandle(datePad(substr(birthday, 4, 6), "month"), "month")
	day := datePipelineHandle(datePad(substr(birthday, 6, 8), "day"), "day")

	addressCodeInt, _ := strconv.Atoi(addressCode)
	if _, ok := data.AddressCodeTimeline[addressCodeInt]; ok {
		timeLine := data.AddressCodeTimeline[addressCodeInt]
		for _, val := range timeLine {
			if val["address"] == address {
				if val["start_year"] != "" {
					startYaer = val["start_year"]
				}
				if val["end_year"] != "" {
					endYear = val["end_year"]
				}
			}
		}
	}

	yearInt, _ := strconv.Atoi(year)
	startYaerInt, _ := strconv.Atoi(startYaer)
	endYearInt, _ := strconv.Atoi(endYear)
	if yearInt < startYaerInt {
		year = startYaer
	}
	if yearInt > endYearInt {
		year = endYear
	}

	birthday = year + month + day
	_, err := time.Parse("20060102", birthday)
	// example: 195578
	if err != nil {
		year = datePad(year, "year")
		month = datePad(month, "month")
		day = datePad(day, "day")
	}

	return year + month + day
}

// 日期处理
func datePipelineHandle(date string, category string) string {
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
func generatorOrderCode(sex int) string {
	rand.Seed(time.Now().Unix())
	orderCode := rand.Intn(999-111) + 111
	if sex != orderCode%2 {
		orderCode--
	}

	return strconv.Itoa(orderCode)
}

// 日期补全
func datePad(date string, category string) string {
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
