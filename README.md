<div align="center" style="padding: 30px 0;">
    <img src="logo.png" width="240">
    <p>Youdu Go SDK</p>
    <p>
        <a target="_blank" href="https://github.com/addcnos/youdu/actions/workflows/lint.yml"><img src="https://github.com/addcnos/youdu/actions/workflows/lint.yml/badge.svg" alt="Lint"></a>
        <a target="_blank" href="https://pkg.go.dev/github.com/addcnos/youdu/v2"><img src="https://pkg.go.dev/badge/github.com/addcnos/youdu/v2" alt="GoDoc"></a>
        <a target="_blank" href="https://goreportcard.com/report/github.com/addcnos/youdu"><img src="https://goreportcard.com/badge/github.com/addcnos/youdu" alt="Go Report Card"></a>
        <img src="https://img.shields.io/badge/Language-Golang-blue.svg" alt="Language">
        <a target="_blank" href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-brightgreen.svg" alt="MIT license"></a>
    </p>
    <p>🚀 Youdu SDK is a Go package that provides API implementation related to <a href="https://youdu.im/doc/api/c01_00002.html" target="_blank">Youdu IM</a></p>
</div>


## Getting Started

[Features](./feature.md)

### Installation

```bash
go get github.com/addcnos/youdu/v2
```
   
### Usage

Please refer to the [Documentation](./docs)

```go
package main

import (
	"context"
	"fmt"

	"github.com/addcnos/youdu/v2"
)

func main() {
	client := youdu.NewClient(&youdu.Config{
		Addr:   "http://examaple",
		Buin:   111222333,
		AppId:  "111222333",
		AesKey: "111333445",
	})

	resp, err := client.SendTextMessage(context.Background(), youdu.TextMessageRequest{
		ToUser:  "11111",
		MsgType: youdu.MsgTypeText,
		Text: youdu.MessageText{
			Content: "hello",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
```

## Contributing

Very welcome to join us! Raise an [Issue](https://github.com/addcnos/youdu/issues/new) or submit a [Pull Request](https://github.com/addcnos/youdu/compare).

[![](https://contributors-img.web.app/image?repo=addcnos/youdu)](https://github.com/addcnos/youdu/graphs/contributors)

## License

[MIT License](LICENSE) © 2022-2023 addcnos