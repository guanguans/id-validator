// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"testing"
)

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestIsValid(t *testing.T) {
	errIds := []string{
		"44030819990110",     // 号码位数不合法
		"111111199901101512", // 地址码不合法
		"440308199902301512", // 出生日期码不合法
		"440308199901101513", // 验证码不合法
		"610104620932690",    // 出生日期码不合法
		"11010119900307867X", // 校验位不合法
		"TES12345678901 j",   // 特殊字符格式不合法
	}
	for _, id := range errIds {
		if IsValid(id, false) {
			t.Errorf("ID must be invalid.: %s", id)
		}
	}

	ids := []string{
		"110101199003078670",
		"440308199901101512",
		"500154199804106120",
		"610104620927690",
		"810000199408230021", // 港澳居民居住证 18 位
		"830000199201300022", // 台湾居民居住证 18 位
		"44040119580101000X", // 历史遗留数据：珠海市市辖区
		"140120197901010008", // 历史遗留数据：太原市市区
		"441282198101011230", // 历史遗留数据：广东省肇庆市罗定市

		"500154199301135886", // 出生日期在地址码发布之前(非严格模式)
		"411082198901010002", // 出生日期在地址码发布之前(非严格模式)
		"370620199505100123", // 出生日期在地址码发布之前(非严格模式)
	}
	for _, id := range ids {
		if !IsValid(id, false) {
			t.Errorf("ID must be valid.: %s", id)
		}
	}

	strictIds := []string{
		"500154199301135886", // 出生日期在地址码发布之前(严格模式)
		"500154199301135886", // 出生日期在地址码发布之前(严格模式)
		"370620199505100123", // 出生日期在地址码发布之前(严格模式)
	}
	for _, id := range strictIds {
		if IsValid(id, true) {
			t.Errorf("ID must be strict valid.: %s", id)
		}
	}
}

func TestGetInfo(t *testing.T) {
	_, e1 := GetInfo("500154199301135886", false)
	if e1 != nil {
		t.Errorf("`e1` must be nil.: %v", e1)
	}

	_, e2 := GetInfo("500154199301135886", true)
	if e2 == nil {
		t.Errorf("`e2` must not be nil.: %v", e2)
	}
}

func TestFakeId(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id := FakeId()
		if !IsValid(id, false) {
			t.Errorf("ID must be valid.: %s", id)
		}
		if l := len(id); l != 18 {
			t.Errorf("ID length must be 15.: %d", l)
		}
	}
}

func TestFakeRequireId(t *testing.T) {
	got1 := IsValid(FakeRequireId(true, "上海市", "2000", 1), false)
	if !got1 {
		t.Errorf("got1 must be true.: %v", got1)
	}

	got2 := IsValid(FakeRequireId(true, "黄浦区", "2001", 0), false)
	if !got2 {
		t.Errorf("got2 must be true.: %v", got2)
	}

	got3 := IsValid(FakeRequireId(true, "江苏省", "200001", 1), false)
	if !got3 {
		t.Errorf("got3 must be true.: %v", got3)
	}

	got4 := IsValid(FakeRequireId(true, "南京市", "2002", 0), false)
	if !got4 {
		t.Errorf("got4 must be true.: %v", got4)
	}

	got5 := IsValid(FakeRequireId(true, "秦淮区", "2003", 0), false)
	if !got5 {
		t.Errorf("got5 must be true.: %v", got5)
	}

	got6 := IsValid(FakeRequireId(true, "台湾省", "20181010", 0), false)
	if !got6 {
		t.Errorf("got6 must be true.: %v", got6)
	}

	got7 := IsValid(FakeRequireId(true, "香港特别行政区", "20181010", 1), false)
	if !got7 {
		t.Errorf("got7 must be true.: %v", got7)
	}

	got8 := IsValid(FakeRequireId(true, "澳门特别行政区", "20181111", 0), false)
	if !got8 {
		t.Errorf("got8 must be true.: %v", got8)
	}

	id := FakeRequireId(false, "江苏省", "19951102", 0)
	if !IsValid(id, false) {
		t.Errorf("ID must be valid.: %s", id)
	}
	if l := len(id); l != 15 {
		t.Errorf("ID length must be 15.: %d", l)
	}
	info, _ := GetInfo(id, false)
	if info.Sex != 0 {
		t.Errorf("Sex must be 0.: %d", info.Sex)
	}
	if info.AddressTree[0] != "江苏省" {
		t.Errorf("Province must be 江苏省.: %s", info.AddressTree[0])
	}
	if birthday := info.Birthday.Format("20060102"); birthday != "19951102" {
		t.Errorf("Birthday must be 19951102.: %s", birthday)
	}
}

func TestUpgradeId(t *testing.T) {
	_, e1 := UpgradeId("610104620927690")
	if e1 != nil {
		t.Errorf("`e1` must be nil.: %v", e1)
	}

	_, e2 := UpgradeId("61010462092769")
	if e2 == nil {
		t.Errorf("`e2` must not be nil.: %v", e2)
	}
}

// TestHistoricalAreaCodes tests validation of IDs with historical area codes
// This addresses issue #60: 行政数据里不含老数据
func TestHistoricalAreaCodes(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		strict    bool
		wantValid bool
		wantAddr  string
	}{
		{
			name:      "370284 born 2000 (valid period) - strict",
			id:        "370284200001010015", // checksum calculated
			strict:    true,
			wantValid: true,
			wantAddr:  "山东省青岛市胶南市",
		},
		{
			name:      "370284 born 2000 (valid period) - loose",
			id:        "370284200001010015",
			strict:    false,
			wantValid: true,
			wantAddr:  "山东省青岛市胶南市",
		},
		{
			name:      "370284 born 1994 (before valid period) - strict",
			id:        "370284199401010019",
			strict:    true,
			wantValid: false,
		},
		{
			name:      "370284 born 1994 (before valid period) - loose",
			id:        "370284199401010019",
			strict:    false,
			wantValid: true,
			wantAddr:  "山东省青岛市胶南市",
		},
		{
			name:      "370284 born 2012 (after valid period) - strict",
			id:        "370284201201010014",
			strict:    true,
			wantValid: false,
		},
		{
			name:      "370284 born 2012 (after valid period) - loose",
			id:        "370284201201010014",
			strict:    false,
			wantValid: true,
			wantAddr:  "山东省青岛市胶南市",
		},
		{
			name:      "370211 current area code - strict",
			id:        "370211200001010017",
			strict:    true,
			wantValid: true,
			wantAddr:  "山东省青岛市黄岛区",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation
			got := IsValid(tt.id, tt.strict)
			if got != tt.wantValid {
				t.Errorf("IsValid() = %v, want %v", got, tt.wantValid)
			}

			// Test GetInfo for valid cases
			if tt.wantValid {
				info, err := GetInfo(tt.id, tt.strict)
				if err != nil {
					t.Errorf("GetInfo() error = %v, want nil", err)
					return
				}

				if info.Address != tt.wantAddr {
					t.Errorf("GetInfo().Address = %v, want %v", info.Address, tt.wantAddr)
				}

				// Verify abandoned status for historical codes
				if tt.id[0:6] == "370284" {
					if info.Abandoned != 1 {
						t.Errorf("Expected historical area code to be marked as abandoned, got %d", info.Abandoned)
					}
				}
			}
		})
	}
}

// TestAreaCodeTransition tests the specific transition mentioned in issue #60
func TestAreaCodeTransition(t *testing.T) {
	// Test case: Someone born in 2000 when 370284 was valid
	// Their ID should validate successfully
	historicalId := "370284200001010015"
	
	// Should be valid in both modes
	if !IsValid(historicalId, false) {
		t.Error("Historical ID should be valid in loose mode")
	}
	
	if !IsValid(historicalId, true) {
		t.Error("Historical ID should be valid in strict mode when birth year is within valid period")
	}
	
	// Should provide proper address information
	info, err := GetInfo(historicalId, false)
	if err != nil {
		t.Fatalf("GetInfo() error = %v", err)
	}
	
	expectedAddr := "山东省青岛市胶南市"
	if info.Address != expectedAddr {
		t.Errorf("Expected address %s, got %s", expectedAddr, info.Address)
	}
	
	// Should be marked as abandoned since area code changed
	if info.Abandoned != 1 {
		t.Errorf("Expected Abandoned=1 for historical area code, got %d", info.Abandoned)
	}
	
	// Birth year should be correctly parsed
	expectedYear := 2000
	if info.Birthday.Year() != expectedYear {
		t.Errorf("Expected birth year %d, got %d", expectedYear, info.Birthday.Year())
	}
}

// TestCurrentAreaCode tests the current area code that replaced the historical one
func TestCurrentAreaCode(t *testing.T) {
	// Test the current area code 370211 (黄岛区)
	currentId := "370211200001010017"
	
	if !IsValid(currentId, true) {
		t.Error("Current area code should be valid in strict mode")
	}
	
	info, err := GetInfo(currentId, true)
	if err != nil {
		t.Fatalf("GetInfo() error = %v", err)
	}
	
	expectedAddr := "山东省青岛市黄岛区"
	if info.Address != expectedAddr {
		t.Errorf("Expected address %s, got %s", expectedAddr, info.Address)
	}
	
	// Should not be marked as abandoned
	if info.Abandoned != 0 {
		t.Errorf("Expected Abandoned=0 for current area code, got %d", info.Abandoned)
	}
}
