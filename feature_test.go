// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import "testing"

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestFeature(t *testing.T) {
	for i := 0; i < 100; i++ {
		if got1 := IsValid(FakeId(), false); !got1 {
			t.Errorf("`got1` must be true.: %v", got1)
		}

		if got2 := IsValid(FakeRequireId(true, "江苏省", "200001", 1), false); !got2 {
			t.Errorf("`got2` must be true.: %v", got2)
		}

		if _, e1 := GetInfo(FakeId(), false); e1 != nil {
			t.Errorf("`e1` must be nil.: %v", e1)
		}

		if _, e2 := GetInfo(FakeRequireId(true, "江苏省", "200001", 1), false); e2 != nil {
			t.Errorf("`e2` must be nil.: %v", e2)
		}

		// id, e3 := UpgradeId(FakeRequireId(false, "", "", 0))
		id, e3 := UpgradeId("610104620927690")
		if e3 != nil {
			t.Errorf("`e3` must be nil.: %v", e3)
		}
		if l := len(id); l != 18 {
			t.Errorf("`id` of length must be 18.: %d", l)
		}
	}
}
