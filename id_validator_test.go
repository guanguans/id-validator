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
	_, err := GetInfo("500154199301135886", false)
	if err != nil {
		t.Errorf("Errors must be nil.")
	}
	_, e := GetInfo("440308199901101513", true)
	if e == nil {
		t.Errorf("Errors must not be nil.")
	}
}

func TestUpgradeId(t *testing.T) {
	_, err := UpgradeId("610104620927690")
	if err != nil {
		t.Errorf("Errors must be nil.")
	}

	_, e := UpgradeId("61010462092769")
	if e == nil {
		t.Errorf("Errors must not be nil.")
	}
}

func TestFakeId(t *testing.T) {
	id := FakeId()
	if len(id) != 18 {
		t.Errorf("String length must be 18. : %s", id)
	}
	if !IsValid(id, false) {
		t.Errorf("%s must be true.", id)
	}
}

func TestFakeRequireId(t *testing.T) {
	id := FakeRequireId(false, "", "", 0)
	if len(id) != 15 {
		t.Errorf("String length must be 15. : %s", id)
	}
	if !IsValid(id, false) {
		t.Errorf("%s must be true.", id)
	}

	info, _ := GetInfo(id, false)
	if info.Sex != 0 {
		t.Errorf("%s must be 0.", "0")
	}
}
