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
	ids := [4]string{
		"500154199301135886",
		"610104620927690",
		"810000199408230021",
		"830000199201300022",
	}
	for _, id := range ids {
		if !IsValid(id, false) {
			t.Errorf("%s must be true.", id)
		}
	}

	errIds := [6]string{
		"440308199901101513",
		"4403081999011015133",
		"510104621927691",
		"61010462092769",
		"810000199408230022",
		"830000199201300023",
	}
	for _, id := range errIds {
		if IsValid(id, true) {
			t.Errorf("%s must be false.", id)
		}
	}
}

func TestGetInfo(t *testing.T) {
	_, e1 := GetInfo("500154199301135886", false)
	if e1 != nil {
		t.Errorf("`e1` must be nil.")
	}

	_, e2 := GetInfo("500154199301135886", true)
	if e2 == nil {
		t.Errorf("`e2` must not be nil.")
	}
}

func TestUpgradeId(t *testing.T) {
	_, e1 := UpgradeId("610104620927690")
	if e1 != nil {
		t.Errorf("`e1` must be nil.")
	}

	_, e2 := UpgradeId("61010462092769")
	if e2 == nil {
		t.Errorf("`e2` must not be nil.")
	}
}

func TestFakeId(t *testing.T) {
	id := FakeId()
	if !IsValid(id, false) {
		t.Errorf("%s must be valid.", id)
	}
	if len(id) != 18 {
		t.Errorf("String length must be 18. : %s", id)
	}
}

func TestFakeRequireId(t *testing.T) {
	id := FakeRequireId(false, "江苏省", "1995", 0)

	if !IsValid(id, false) {
		t.Errorf("%s must be valid.", id)
	}

	if len(id) != 15 {
		t.Errorf("String length must be 15. : %s", id)
	}

	info, _ := GetInfo(id, false)
	if info.Sex != 0 {
		t.Errorf("`Sex` must be 0.")
	}

	if info.AddressTree[0] != "江苏省" {
		t.Errorf("`province` must be `江苏省`.")
	}

	if info.Birthday.Format("2006") != "1995" {
		t.Errorf("Year of birth must be `1995`.")
	}
}
