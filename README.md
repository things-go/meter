# meter storage metering
storage metering,like B,KB,MB,GB,TB,PB,EB

[![GoDoc](https://godoc.org/github.com/thinkgos/meter?status.svg)](https://godoc.org/github.com/thinkgos/meter)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/meter?tab=doc)
[![Build Status](https://www.travis-ci.org/thinkgos/meter.svg?branch=master)](https://www.travis-ci.org/thinkgos/meter)
[![codecov](https://codecov.io/gh/thinkgos/meter/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/meter)
![Action Status](https://github.com/thinkgos/meter/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/meter)](https://goreportcard.com/report/github.com/thinkgos/meter)
[![License](https://img.shields.io/github/license/thinkgos/meter)](https://github.com/thinkgos/meter/raw/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/meter)](https://github.com/thinkgos/meter/tags)


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
    go get -u github.com/thinkgos/meter
```

2. Import it in your code.

```go
    import "github.com/thinkgos/meter"
```

## Quick start

```go
    package main
    
    import (
        "fmt"

        "github.com/thinkgos/meter"
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