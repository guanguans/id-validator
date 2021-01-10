package id_validator

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
