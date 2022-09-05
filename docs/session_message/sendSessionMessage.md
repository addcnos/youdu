[查看所有](./README.md)

### 发送消息

参考 [有度即时通讯服务端 API 文档 -- 会话消息 -- 发送消息](https://youdu.im/doc/api/c01_00009.html)

#### 消息类型
会话消息共有以下分类
1. [文本消息](#文本消息)  `session.TextMessage`
2. [图片消息](#图片消息)  `message.ImageMessage`
3. [文件消息](#文件消息)  `message.FileMessage `
4. [语音消息](#语音消息)  `message.VoiceMessage`
5. [视频消息](#视频消息)  `message.VideoMessage`

#### 发送消息

消息的发送很简单，只需完成先创建消息对象，就可以调用方法进行发送
```go
import 	"github.com/addcnos/youdu/session"
// 发送
yd.Session().Send(msg session.Message)
```

##### 文本消息

```go
import "github.com/addcnos/youdu/session"

TextMessage := &session.TextMessage{
    SessionId: "sessionId",//发送给多人会话时填写
    Receiver:  "user1",//发送给单人会话时填写
    Sender:    "sender",//消息发送者账号
    MsgType:   session.MsgTypeText
    Text: &session.TextItem{
        Content: "content",//消息内容，支持表情
    },
}
yd.Session().Send(TextMessage)
```
#####  图片消息

```go
import "github.com/addcnos/youdu/session"
imageMessage:= &session.ImageMessage{
    SessionId: "sessionId",//发送给多人会话时填写
    Receiver:  "user1",//发送给单人会话时填写
    Sender:    "sender",//消息发送者账号
    MsgType:   session.MsgTypeImage,
    Image: &session.MediaItem{
        MediaId: "mediaId",//图片素材文件id。通过上传素材文件接口获取
    },

}
yd.Session().Send(imageMessage)
``` 
#####  文件消息

```go
import "github.com/addcnos/youdu/session"
fileMessage:= &session.FileMessage{
    SessionId: "sessionId",//发送给多人会话时填写
    Receiver:  "user1",//发送给单人会话时填写
    Sender:    "sender",//消息发送者账号
    MsgType:   session.MsgTypeFile,
    File: &session.MediaItem{
        MediaId: "mediaId",//素材文件id。通过上传素材文件接口获取
    },

}
yd.Session().Send(fileMessage)
``` 
#####  语音消息

```go
import "github.com/addcnos/youdu/session"
voiceMessage:= &session.VoiceMessage{
    SessionId: "sessionId",//发送给多人会话时填写
    Receiver:  "user1",//发送给单人会话时填写
    Sender:    "sender",//消息发送者账号
    MsgType:   session.MsgTypeVoice,
    Voice: &session.MediaItem{
        MediaId: "mediaId",//素材文件id。通过上传素材文件接口获取
    },

}
yd.Session().Send(voiceMessage)
``` 

#####  视频消息

```go
import "github.com/addcnos/youdu/session"
videoMessage:= &session.VideoMessage{
    SessionId: "sessionId",//发送给多人会话时填写
    Receiver:  "user1",//发送给单人会话时填写
    Sender:    "sender",//消息发送者账号
    MsgType:   session.MsgTypeVideo,
    Video: &session.MediaItem{
        MediaId: "mediaId",//素材文件id。通过上传素材文件接口获取
    },

}
yd.Session().Send(videoMessage)
``` 