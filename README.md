<p align="center">
goassist is provides utility apis for gopher more easily to use.
</p>

[![Go Report Card](https://goreportcard.com/badge/github.com/jhunters/goassist)](https://goreportcard.com/report/github.com/jhunters/goassist)
[![Build Status](https://github.com/jhunters/goassist/actions/workflows/go.yml/badge.svg)](https://github.com/jhunters/goassist/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/jhunters/goassist/branch/main/graph/badge.svg)](https://codecov.io/gh/jhunters/goassist)
[![Releases](https://img.shields.io/github/release/jhunters/goassist/all.svg?style=flat-square)](https://github.com/jhunters/goassist/releases)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/jhunters/goassist)
[![LICENSE](https://img.shields.io/github/license/jhunters/goassist.svg?style=flat-square)](https://github.com/jhunters/goassist/blob/main/LICENSE)


# Go Required Version
need go 1.21

# Install

go get github.com/jhunters/goassist

# API List

API包|说明|文档
--|--|--
arrayutil|数组处理|[doc](https://pkg.go.dev/github.com/jhunters/goassist/arrayutil)
concurrent|并发操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/concurrent)
concurrent/syncx| 并发同步应用(channel, pool, map)|[doc](https://pkg.go.dev/github.com/jhunters/goassist/concurrent/syncx)
concurrent/atomicx|原子操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/concurrent/actomicx)
containerx|容器操作 | [heap](https://pkg.go.dev/github.com/jhunters/goassist/container/heapx) [list](https://pkg.go.dev/github.com/jhunters/goassist/container/listx) [map](https://pkg.go.dev/github.com/jhunters/goassist/container/mapx) [queue](https://pkg.go.dev/github.com/jhunters/goassist/container/queue) [ring](https://pkg.go.dev/github.com/jhunters/goassist/container/ringx) [set](https://pkg.go.dev/github.com/jhunters/goassist/container/set) [stack](https://pkg.go.dev/github.com/jhunters/goassist/container/stack)
hashx|hash操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/hashx)
maputil|map操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/maputil)
reflectutil|反射操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/reflectutil)
stringutil|字符串操作|[doc](https://pkg.go.dev/github.com/jhunters/goassist/stringutil)
unsafex|unsafe包扩展应用|[doc](https://pkg.go.dev/github.com/jhunters/goassist/unsafex)
timeutil|时间处理|[doc](https://pkg.go.dev/github.com/jhunters/goassist/timeutil)
web|http文件处理|[doc](https://pkg.go.dev/github.com/jhunters/goassist/web)
netutil|net工具类|[doc](https://pkg.go.dev/github.com/jhunters/goassist/netutil)

## License
goassist is [Apache 2.0 licensed](./LICENSE).