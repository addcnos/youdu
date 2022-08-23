# Youdu SDK

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

Youdu SDK 是一个 Go 包，它提供了[有度即时通讯](https://youdu.im/doc/api/c01_00002.html) 服务端 Api 的 Golang 实现。

## 内容列表
- [Youdu SDK](#youdu-sdk)
  - [内容列表](#内容列表)
  - [依赖](#依赖)
  - [安装](#安装)
  - [使用说明](#使用说明)
  - [详细文档](#详细文档)
    - [生成器](#生成器)
  - [如何贡献](#如何贡献)
  - [License](#license)

## 依赖
- Golang 版本 1.18
- 配置环境变量 `go env -w GO111MODULE=on`

## 安装
1. 首先确保您已经安装了 golang 以及正确的设置了 GO111MODULE，然后您可以使用以下命令将 `youdu-sdk` 作为依赖添加到您的 Go 程序中。 
    ```go
      go get -u github.com/addcnos/youdu
    ```
2. 在您的代码中导入它：
    ```go
      import "github.com/addcnos/youdu"
    ```
## 使用说明
1. 首先需要实例化 youdu ，后续所有的操作均依赖此实例
    ```go
      import "github.com/addcnos/youdu"

      yd := youdu.New(&youdu.Config{
        Api:    "http://domain.com/api", //youdu api host
        Buin:   1111111, //企业 buin 码
        AppId:  "22222222222222", //应用appId
        AesKey: "3444444444444444444444444444444444", //应用 AesKey
      })
    ```
2. 以发送应用消息为例：
   发送文本消息 
    ```go
    yd.Message().Send(&message.TextMessage{
        ToUser:  "11111",//指定用户
        ToDept:  "",//指定部门
        MsgType: message.MsgTypeText,
        Text: &message.TextItem{
          Content: "test",
        }
      })
    ```
## 详细文档
关于具体的使用方法，请参照以下文档
1. [企业应用](docs/EnterpriseApplication/enterpriseMessage.md) 
      - [应用消息](docs/EnterpriseApplication/enterpriseMessage.md) 
      - [弹窗消息](docs/EnterpriseApplication/enterpriseMessage.md) 
2. [单点登录](docs/EnterpriseApplication/enterpriseMessage.md) 
3. [会话管理](docs/EnterpriseApplication/enterpriseMessage.md)
4. [会话消息](docs/EnterpriseApplication/enterpriseMessage.md)
      - [发送消息](docs/EnterpriseApplication/enterpriseMessage.md)
5. [企业通讯录](docs/paragraph.md)
      - [部门管理](docs/EnterpriseApplication/enterpriseMessage.md)
      - [用户管理](docs/EnterpriseApplication/enterpriseMessage.md)
      - [第三方认证](docs/EnterpriseApplication/enterpriseMessage.md)
      - [群管理](docs/EnterpriseApplication/enterpriseMessage.md)
6. [素材管理](docs/number.md)
### 生成器

想要使用生成器的话，请看 [generator-standard-readme](https://github.com/RichardLitt/generator-standard-readme)。
有一个全局的可执行文件来运行包里的生成器，生成器的别名叫 `standard-readme`。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/RichardLitt/standard-readme/issues/new) 或者提交一个 Pull Request。


标准 Readme 遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。

##  License

[MIT](LICENSE) © Richard Littauer