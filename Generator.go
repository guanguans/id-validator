package id_validator

import (
	"math"
	"strconv"
	"strings"
)

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
