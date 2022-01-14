// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import "testing"

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestFeature(t *testing.T) {
	isValid1 := IsValid(FakeId(), false)
	if !isValid1 {
		t.Errorf("`isValid1` must be true.")
	}

	isValid2 := IsValid(FakeRequireId(true, "江苏省", "200001", 1), true)
	if !isValid2 {
		t.Errorf("`isValid2` must be true.")
	}

	_, err1 := GetInfo(FakeRequireId(true, "江苏省", "200001", 1), true)
	if err1 != nil {
		t.Errorf("`err1` must be nil.")
	}

	_, err2 := GetInfo(FakeRequireId(true, "江苏省", "200001", 1), true)
	if err2 != nil {
		t.Errorf("`err2` must be nil.")
	}

	upgradedId, err3 := UpgradeId("610104620927690")
	if err3 != nil {
		t.Errorf("`err3` must be nil.")
	}
	if len(upgradedId) != 18 {
		t.Errorf("`upgradedId`  length must be 18.:%d", len(upgradedId))
	}
}
