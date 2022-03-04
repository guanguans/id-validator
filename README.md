# id-validator

[简体中文](README.md) | [ENGLISH](README-EN.md)

> 中国身份证号验证器。

[![tests](https://github.com/guanguans/id-validator/actions/workflows/tests.yml/badge.svg)](https://github.com/guanguans/id-validator/actions/workflows/tests.yml)
[![gocover.io](https://gocover.io/_badge/github.com/guanguans/id-validator)](https://gocover.io/github.com/guanguans/id-validator)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/id-validator)](https://goreportcard.com/report/github.com/guanguans/id-validator)
[![GoDoc](https://godoc.org/github.com/guanguans/id-validator?status.svg)](https://godoc.org/github.com/guanguans/id-validator)
[![GitHub release](https://img.shields.io/github/tag/guanguans/id-validator.svg)](https://github.com/guanguans/id-validator/releases)
[![GitHub license](https://img.shields.io/github/license/guanguans/id-validator.svg)](https://github.com/guanguans/id-validator/blob/master/LICENSE)

## 功能

* 验证中国身份证号
* 获取身份证号信息
* 升级 15 位身份证号为 18 位
* 伪造符合校验的身份证号

## 环境要求

* Go >= 1.11

## 安装

```shell script
$ go get -u github.com/guanguans/id-validator
```

## 使用

这只是一个快速介绍, 请查看 [GoDoc](https://godoc.org/github.com/guanguans/id-validator) 获得详细信息。

```go
package main

import (
    idvalidator "github.com/guanguans/id-validator"
    "gopkg.in/ffmt.v1"
)

func main() {
    // 验证身份证号合法性
    ffmt.P(idvalidator.IsValid("500154199301135886", true))  // 严格模式验证大陆居民身份证18位
    ffmt.P(idvalidator.IsValid("500154199301135886", false)) // 非严格模式验证大陆居民身份证18位
    ffmt.P(idvalidator.IsValid("11010119900307803X", false)) // 大陆居民身份证末位是X18位
    ffmt.P(idvalidator.IsValid("610104620927690", false))    // 大陆居民身份证15位
    ffmt.P(idvalidator.IsValid("810000199408230021", false)) // 港澳居民居住证18位
    ffmt.P(idvalidator.IsValid("830000199201300022", false)) // 台湾居民居住证18位

    // 获取身份证号信息
    ffmt.P(idvalidator.GetInfo("500154199301135886", true))  // 严格模式获取身份证号信息
    ffmt.P(idvalidator.GetInfo("500154199301135886", false)) // 非严格模式获取身份证号信息
    // []interface {}[
    //     github.com/guanguans/id-validator.IdInfo{          // 身份证号信息
    //         AddressCode: int(500154)                           // 地址码
    //         Abandoned:   int(0)                                // 地址码是否废弃：1为废弃的，0为正在使用的
    //         Address:     string("重庆市市辖区开州区")             // 地址
    //         AddressTree: []string[                             // 省市区三级列表
    //             string("重庆市")                                    // 省
    //             string("市辖区")                                    // 市
    //             string("开州区")                                    // 区
    //         ]
    //         Birthday:      <1993-01-13 00:00:00 +0800 CST>     // 出生日期
    //         Constellation: string("摩羯座")                     // 星座
    //         ChineseZodiac: string("酉鸡")                       // 生肖
    //         Sex:           int(0)                              // 性别：1为男性，0为女性
    //         Length:        int(18)                             // 号码长度
    //         CheckBit:      string("6")                         // 校验码
    //     }
    //     <nil>                                              // 错误信息
    // ]

    // 生成可通过校验的假身份证号
    ffmt.P(idvalidator.FakeId())                                  // 随机生成
    ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1)) // 生成出生于2000年1月江苏省的男性居民身份证

    // 15位号码升级为18位
    ffmt.P(idvalidator.UpgradeId("610104620927690"))
    // []interface {}[
    // 	string("610104196209276908") // 升级后号码
    // 	<nil>                        // 错误信息
    // ]
}
```

## 测试

```shell script
$ make test
```

## 变更日志

请参阅 [CHANGELOG](CHANGELOG.md) 获取最近有关更改的更多信息。

## 贡献指南

请参阅 [CONTRIBUTING](.github/CONTRIBUTING.md) 有关详细信息。

## 安全漏洞

请查看[我们的安全政策](../../security/policy)了解如何报告安全漏洞。

## 贡献者

* [guanguans](https://github.com/guanguans)
* [所有贡献者](../../contributors)

## 相关项目

* [jxlwqq/id-validator](https://github.com/jxlwqq/id-validator)，jxlwqq
* [jxlwqq/id-validator.py](https://github.com/jxlwqq/id-validator.py)，jxlwqq
* [mc-zone/IDValidator](https://github.com/mc-zone/IDValidator)，mc-zone
* [renyijiu/id_validator](https://github.com/renyijiu/id_validator)，renyijiu

## 参考资料

* [中华人民共和国公民身份号码](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
* [中华人民共和国民政部：行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)
* [中华人民共和国行政区划代码历史数据集](https://github.com/jxlwqq/address-code-of-china)
* [国务院办公厅关于印发《港澳台居民居住证申领发放办法》的通知](http://www.gov.cn/zhengce/content/2018-08/19/content_5314865.htm)
* [港澳台居民居住证](https://zh.wikipedia.org/wiki/港澳台居民居住证)

## 协议

MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。
