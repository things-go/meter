# meter storage metering
storage metering,like B,KB,MB,GB,TB,PB,EB

[![GoDoc](https://godoc.org/github.com/thinkgos/meter?status.svg)](https://godoc.org/github.com/thinkgos/meter)
[![Build Status](https://www.travis-ci.org/thinkgos/meter.svg?branch=master)](https://www.travis-ci.org/thinkgos/meter)
[![codecov](https://codecov.io/gh/thinkgos/meter/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/meter)
![Action Status](https://github.com/thinkgos/meter/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/meter)](https://goreportcard.com/report/github.com/thinkgos/meter)

## Features
1. support human string like B,KB,MB,GB,TB,PB,EB
2. support parse string like:
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
        s := meter.HumanSize(100) 
        value, _ := meter.ParseBytes("100k")
        fmt.Println(s)
        fmt.Println(value)

        // output:
        // 100.0B
        // 102400
    }
```

## Reference