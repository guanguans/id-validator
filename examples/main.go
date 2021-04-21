package main

import (
	idvalidator "github.com/guanguans/id-validator"
	"gopkg.in/ffmt.v1"
)

func main() {
	// 验证身份证号合法性
	ffmt.P(idvalidator.IsValid("440308199901101512")) // 大陆居民身份证18位
	ffmt.P(idvalidator.IsValid("11010119900307803X")) // 大陆居民身份证末位是X18位
	ffmt.P(idvalidator.IsValid("610104620927690"))    // 大陆居民身份证15位
	ffmt.P(idvalidator.IsValid("810000199408230021")) // 港澳居民居住证18位
	ffmt.P(idvalidator.IsValid("830000199201300022")) // 台湾居民居住证18位

	// 获取身份证号信息
	ffmt.P(idvalidator.GetInfo("440308199901101512"))

	// 生成可通过校验的假身份证号
	ffmt.P(idvalidator.FakeId())                                // 随机生成
	ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1)) // 生成出生于2000年1月江苏省的男性居民身份证

	// 15位号码升级为18位
	ffmt.P(idvalidator.UpgradeId("610104620927690"))
}
