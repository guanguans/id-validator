package idvalidator

import (
	"testing"
)

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestIsValidBasic(t *testing.T) {
	errIds := []string{
		"44030819990110",     // 号码位数不合法
		"440308199902301512", // 出生日期码不合法
		"440308199901101513", // 验证码不合法
		"610104620932690",    // 出生日期码不合法
		"11010119900307867X", // 校验位不合法
		"TES12345678901 j",   // 特殊字符格式不合法
	}
	for _, id := range errIds {
		if IsValidBasic(id) {
			t.Errorf("ID must be invalid.: %s", id)
		}
	}
}

func TestGetInfoBasic(t *testing.T) {
	_, e1 := GetBasicInfo("440301197110292910")
	if e1 != nil {
		t.Errorf("`e1` must be nil.: %v", e1)
	}

	_, e2 := GetBasicInfo("500154199302305886")
	if e2 == nil {
		t.Errorf("`e2` must not be nil.: %v", e2)
	}
}
