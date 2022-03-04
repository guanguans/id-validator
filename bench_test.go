// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package idvalidator

import (
	"testing"
)

func BenchmarkIsValid(b *testing.B) {
	benchmarks1 := []struct {
		name string
		id   string
	}{
		{id: "440308199902301512"}, // 无效(出生日期码不合法)
		{id: "11010119900307867X"}, // 无效(校验位不合法)
		{id: "441282198101011230"}, // 特殊(历史遗留数据)
		{id: "370620199505100123"}, // 特殊(出生日期在地址码发布之前)
		{id: "110101199003078670"}, // 正常
	}
	for _, bm := range benchmarks1 {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsValid(bm.id, false)
			}
		})
	}

	benchmarks2 := []struct {
		name string
		id   string
	}{
		{id: "370620199505100123"}, // 特殊(出生日期在地址码发布之前)
	}
	for _, bm := range benchmarks2 {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsValid(bm.id, true)
			}
		})
	}
}

func BenchmarkGetInfo(b *testing.B) {
	benchmarks1 := []struct {
		name string
		id   string
	}{
		{id: "11010119900307867X"}, // 无效(校验位不合法)
		{id: "441282198101011230"}, // 特殊(历史遗留数据：广东省肇庆市罗定市)
		{id: "370620199505100123"}, // 特殊(出生日期在地址码发布之前)
		{id: "110101199003078670"}, // 正常

	}
	for _, bm := range benchmarks1 {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetInfo(bm.id, false)
			}
		})
	}

	benchmarks2 := []struct {
		name string
		id   string
	}{
		{id: "500154199301135886"}, // 特殊(出生日期在地址码发布之前)

	}
	for _, bm := range benchmarks2 {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetInfo(bm.id, true)
			}
		})
	}
}

func BenchmarkFakeId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FakeId()
	}
}

func BenchmarkFakeRequireId(b *testing.B) {
	benchmarks := []struct {
		name       string
		isEighteen bool
		address    string
		birthday   string
		sex        int
	}{
		{isEighteen: false, address: "浙江省", birthday: "20000101", sex: 1},
		{isEighteen: true, address: "台湾省", birthday: "20000101", sex: 1},
		{isEighteen: true, address: "香港特别行政区", birthday: "20000101", sex: 0},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FakeRequireId(bm.isEighteen, bm.address, bm.birthday, bm.sex)
			}
		})
	}
}

func BenchmarkUpgradeId(b *testing.B) {
	benchmarks := []struct {
		name string
		id   string
	}{
		{id: "610104620927690"}, // 有效
		{id: "61010462092769"},  // 无效
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UpgradeId(bm.id)
			}
		})
	}
}
