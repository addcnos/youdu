<div align="center" style="padding: 30px 0;">
    <img src="logo.png" width="240">
    <p>Youdu Go SDK</p>
    <p>
        <a target="_blank" href="https://github.com/addcnos/youdu/actions/workflows/lint.yml"><img src="https://github.com/addcnos/youdu/actions/workflows/lint.yml/badge.svg" alt="Lint"></a>
        <a target="_blank" href="https://pkg.go.dev/github.com/addcnos/youdu"><img src="https://pkg.go.dev/badge/github.com/addcnos/youdu" alt="GoDoc"></a>
        <a target="_blank" href="https://goreportcard.com/report/github.com/addcnos/youdu"><img src="https://goreportcard.com/badge/github.com/addcnos/youdu" alt="Go Report Card"></a>
        <img src="https://img.shields.io/badge/Language-Golang-blue.svg" alt="Language">
        <a target="_blank" href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-brightgreen.svg" alt="MIT license"></a>
    </p>
    <p>🚀 Youdu SDK is a Go package that provides API implementation related to <a href="https://youdu.im/doc/api/c01_00002.html" target="_blank">Youdu IM</a></p>
</div>


## Getting Started

[Features](./feature.md)

### Prerequisites

- Golang Version >= 1.16

### Installation

```bash
go get github.com/addcnos/youdu
```
   
### Usage

Please refer to the [Documentation](./docs)

```go
package main

import (
	"github.com/addcnos/youdu"
	"github.com/addcnos/youdu/message"
)

func main()  {

	yd := youdu.New(&youdu.Config{
		Api:    "http://domain.com/api", // youdu api host
		Buin:   1111111, // 企业 buin 码
		AppId:  "22222222222222", // 应用 appId
		AesKey: "3444444444444444444444444444444444", // 应用 AesKey
	})

	yd.Message().Send(&message.TextMessage{
		ToUser:  "user1|user2", // 指定用户
		ToDept:  "dep1|dep2",   // 指定部门
		MsgType: message.MsgTypeText,
		Text: &message.TextItem{
			Content: "content",
		},
	})
}
```

## Contributing

Very welcome to join us! Raise an [Issue](https://github.com/addcnos/youdu/issues/new) or submit a [Pull Request](https://github.com/addcnos/youdu/compare).

[![](https://contributors-img.web.app/image?repo=addcnos/youdu)](https://github.com/addcnos/youdu/graphs/contributors)

## License

[MIT License](LICENSE) © 2022 addcnos