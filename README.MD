[![Build Status](https://travis-ci.org/akominch/yeelight.svg?branch=master)](https://travis-ci.org/akominch/yeelight)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/akominch/yeelight/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/akominch/yeelight)](https://goreportcard.com/report/github.com/akominch/yeelight)
[![Code Coverage](https://codecov.io/gh/akominch/yeelight/branch/master/graph/badge.svg)](https://codecov.io/gh/akominch/yeelight)

# Yeelight Golang API
Golang API for [Yeelight](https://www.yeelight.com/)

## Overview
Yeelight is simple command line tool and Golang implementation API of Yeelight protocol 
with notification listening support

## Installation
Make sure you have a working Go environment. [See Golang install instructions](http://golang.org/doc/install.html)

To install, run:
```sh
go get github.com/akominch/yeelight
```

## Usage
```go
import (
	"github.com/akominch/yeelight"
	t "github.com/akominch/yeelight/transitions"
)

func main() {
        config := yeelight.YeelightConfig{
		Ip: "192.168.1.24",
		Effect: yeelight.Smooth,
	}
	bulb := yeelight.New(config)

	transitions := t.Alarm()
	flow := yeelight.NewFlow(3, yeelight.Off, transitions)
	_, _ = bulb.StartFlow(flow)
}
```

## API Specification
Yeelight API Specification [can be found here] (https://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf)

## License
yeelight is distributed under the [MIT license](https://opensource.org/licenses/MIT)

## Legal
Yeelight® is a registered trademark of [Yeelight](https://www.yeelight.com/).