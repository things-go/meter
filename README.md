# meter storage metering
storage metering,like B,KB,MB,GB,TB,PB,EB

[![GoDoc](https://godoc.org/github.com/things-go/meter?status.svg)](https://godoc.org/github.com/things-go/meter)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/meter?tab=doc)
[![Build Status](https://www.travis-ci.com/things-go/meter.svg?branch=master)](https://www.travis-ci.com/things-go/meter)
[![codecov](https://codecov.io/gh/things-go/meter/branch/master/graph/badge.svg)](https://codecov.io/gh/things-go/meter)
![Action Status](https://github.com/things-go/meter/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/meter)](https://goreportcard.com/report/github.com/things-go/meter)
[![License](https://img.shields.io/github/license/things-go/meter)](https://github.com/things-go/meter/raw/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/meter)](https://github.com/things-go/meter/tags)


## Features
1. support human string like B,KB,MB,GB,TB,PB,EB
2. support parse string unit below:
    - "empty",b,byte,
    - k,kb,kilo,kilobyte
    - m,mb,mega,megabyte,megabytes
    - g,gb,giga,gigabyte,gigabytes
    - t,tb,tera,terabyte,terabytes
    - p,pb,peta,petabyte,petabytes
    - e,eb,exa,exabyte,exabytes


## Installation

1. You can use the below Go command to install meter.

```bash
    go get -u github.com/things-go/meter
```

2. Import it in your code.

```go
    import "github.com/things-go/meter"
```

## Quick start

[embedmd]:# (_example/main.go go)
```go
package main

import (
	"fmt"

	"github.com/things-go/meter"
)

func main() {
	value, _ := meter.ParseBytes("2.99TB")
	s := meter.HumanSize(value)
	fmt.Println(value)
	fmt.Println(s)

	// output:
	// 3287539767050
	// 2.99TB
}
```

## Reference

## Donation

if package help you a lot,you can support us by:

**Alipay**

![alipay](https://github.com/thinkgos/thinkgos/blob/master/asserts/alipay.jpg)

**WeChat Pay**

![wxpay](https://github.com/thinkgos/thinkgos/blob/master/asserts/wxpay.jpg)
