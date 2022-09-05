[查看所有](./README.md)
### 应用消息

参考 [有度即时通讯服务端 API 文档 -- 企业应用 -- 发送应用消息](https://youdu.im/doc/api/c01_00003.html)
#### 消息类型

应用消息共有以下分类
1. [文本消息](#文本消息)  `message.MsgTypeText`
2. [图片消息](#图片消息)  `message.MsgTypeImage`
3. [文件消息](#文件消息)  `message.MsgTypeFile`
4. [图文消息](#图文消息)  `message.MsgTypeMpNews`
5. [隐式链接](#隐式链接)  `message.MsgTypeLink`
6. [外链消息](#外链消息)  `message.MsgTypeExtLink`
7. [系统消息](#系统消息)  `message.MsgTypeSys`
8. [短信消息](#短信消息)（短信网关收到的手机回复的短信）  `message.MsgTypeSms`
9. [邮件消息](#邮件消息)  `message.MsgTypeMail`

#### 发送消息

消息的发送很简单，只需完成先创建消息对象，就可以调用方法进行发送
```go
import 	"github.com/addcnos/youdu/message"
// 发送
yd.Message().Send(msg message.Message)
```

##### 文本消息

你可以简单发送 
```go
import "github.com/addcnos/youdu/message"

yd.Message().SendText(toUser, content)
```

或者
```go
import "github.com/addcnos/youdu/message"
// 实例文本消息
textMessage := &message.TextMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeText,
    Text: &message.TextItem{
        Content: "hello world!",
    },
}
// 发送
yd.Message().Send(textMessage)
```
#####  图片消息

```go
import "github.com/addcnos/youdu/message"
// 实例图片消息
imageMessage := &message.ImageMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeText,
    Image: &message.MediaItem{
        MediaId: "123", //素材文件ID。通过上传素材文件接口获取
    },
}
// 发送
yd.Message().Send(imageMessage)
``` 
#####  文件消息

```go
import "github.com/addcnos/youdu/message"
// 实例文件消息
fileMessage := &message.FileMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeFile,
    File: &message.MediaItem{
        MediaId: "123", //素材文件ID。通过上传素材文件接口获取
    },
}
// 发送
yd.Message().Send(fileMessage)
``` 
#####  图文消息

```go
import "github.com/addcnos/youdu/message"
// 实例图文消息
mpNewsMessage := &message.MpNewsMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeFile,
    MpNews: []*message.MpNewsItem{
        {
            Title:     "hello world",//标题
            MediaId:   "123",//封面图片ID。通过上传素材文件接口获取
            Content:   "hello world",//正文，最长不超过600个字符，超出部分将自动截取
            Digest:    "hello world",//摘要，最长不超过120个字符，超出部分将自动截取
            ShowFront: 1,//正文是否显示封面图片。1：显示，0：不显示
        },
    },
}
// 发送
yd.Message().Send(mpNewsMessage)
``` 

#####  隐式链接

```go
import "github.com/addcnos/youdu/message"
// 实例隐式链接
linkMessage := &message.MpNewsMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeLink,
    Link: &message.LinkItem{
			Title:  "title",//标题。最多允许64个字符
			Url:    "title",//链接
			Action: 1,//链接打开方式。0：直接打开url；1：url后面带上有度客户端userName和token，可做单点登录
	},
}
// 发送
yd.Message().Send(linkMessage)
``` 

#####  外链消息

```go
import "github.com/addcnos/youdu/message"
// 实例外链消息
exLinkMessage := &message.ExLinkMessage{
    ToUser:  "user1|user2|user3",
    ToDept:  "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeExtLink,
    ExLink: []*message.ExLinkItem{//外链详细信息，最多不超过5条
        {
            Title:   "title",//标题。最多允许64个字符
            Url:     "url",//链接
            Digest:  "hello world",//摘要，最长不超过120个字符，超出部分将自动截取
            MediaId: "123",//封面图片的ID。通过上传素材文件接口获取
        },
    },
}
// 发送
yd.Message().Send(exLinkMessage)
``` 

#####  系统消息

```go
import "github.com/addcnos/youdu/message"
// 实例系统消息
sysMessage := &message.SysMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    ToAll: struct {//需要toAll时，此项必填。true：仅发送给全体在线成员，离线成员即使再上线也不会接收到该消息；false：发送给全体成员，离线成员再上线时也可以收到该消息
        OnlyOnline bool "json:\"onlyOnline\""
    }{
        OnlyOnline: true,
    },
    MsgType: message.MsgTypeSys,
    SysMsg: &message.SysMsgItem{
        Title:       "title",//系统消息标题。最多允许64个字符
        PopDuration: 5,//弹窗显示时间，0及负数弹窗不消失，正数为对应显示秒数
        Msg: []*message.SysMsgItemMsg{//消息详细内容。系统消息内容支持(1)文本和(5)隐式链接
            {
                Text: &message.TextItem{
                    Content: "content",
                },
                Link: &message.LinkItem{
                    Title:  "title",
                    Url:    "url",
                    Action: 1,
                },
            },
        },
    },
}
// 发送
yd.Message().Send(sysMessage)
``` 

#####  短信消息

```go
import "github.com/addcnos/youdu/message"
// 实例短信消息
smsMessage := &message.SmsMessage{
    ToUser:  "user1|user2|user3",
    ToDept:  "deptId1|deptId2|deptId3",
    MsgType: message.MsgTypeSms,
    Sms: &struct{
        From    string `json:"from"` //发送短信的手机号码
        Content string `json:"content"` //消息内容，支持表情，最长不超过600个字符，超出部分将自动截取
    }{
        From:    "from",
        Content: "content",
    },
}
// 发送
yd.Message().Send(smsMessage)
``` 

#####  邮件消息

```go
import "github.com/addcnos/youdu/message"
// 实例邮件消息
mailMessage := &message.MailMessage{
    ToUser:  "user1|user2|user3",
    ToEmail: "email1|email2|email3",//接收成员邮件账号列表。多个接收者用竖线分隔，最多支持1000个。toUser不为空，toEmail值无效
    MsgType: message.MsgTypeMail,
    Mail: &struct {
        Action      string "json:\"action\""//邮件消息类型。new: 新邮件，unread: 未读邮件数
        Subject     string "json:\"subject\""//邮件主题。action为new时有效，可为空
        FromUser    string "json:\"fromUser\""//发送者帐号，action为new时有效
        FromEmail   string "json:\"fromEmail\""//发送者邮件帐号，action为new时有效。fromUser不为空，fromEmail值无效
        Time        int    "json:\"time\""//邮件发送时间。为空默认取服务器接收到消息的时间
        Link        string "json:\"link\""//邮件链接。action为new时有效，点此链接即可打开邮件，为空时点击邮件消息默认执行企业邮箱单点登录
        UnreadCount int    "json:\"unreadCount\""//未读邮件数。action为unread时有效
    }{
        Action:      "cc",
        Subject:     "subject",
        FromUser:    "fromUser",
        FromEmail:   "FromEmail",
        Time:        1542699050,
        Link:        "Link",
        UnreadCount: 10,
    },
}
// 发送
yd.Message().Send(mailMessage)
``` 