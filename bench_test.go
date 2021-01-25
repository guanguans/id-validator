package idvalidator

import (
	"testing"
)

func BenchmarkIsValid(b *testing.B) {
	benchmarks := []struct {
		name string
		id   string
	}{
		{id: "440308199901101512"},
		{id: "610104620927690"},
		{id: "810000199408230021"},
		{id: "830000199201300022"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsValid(bm.name)
			}
		})
	}
}

func BenchmarkGetInfo(b *testing.B) {
	benchmarks := []struct {
		name string
		id   string
	}{
		{id: "440308199901101512"},
		{id: "610104620927690"},
		{id: "810000199408230021"},
		{id: "830000199201300022"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetInfo(bm.name)
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
		{isEighteen: true, address: "浙江省", birthday: "20000101", sex: 0},
		{isEighteen: true, address: "台湾省", birthday: "20000101", sex: 0},
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
		{id: "610104620927690"},
		{id: "61010462092769"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {

			}
		})
	}
}
