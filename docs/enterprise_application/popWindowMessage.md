[查看所有](./README.md)
## 弹窗消息

参考 [有度即时通讯服务端 API 文档 -- 企业应用 -- 发送弹窗消息](https://youdu.im/doc/api/c01_00006.html)

### 发送弹窗消息

```go
//创建弹窗消息
popWindowMessage := &message.PopWindowMessage{
    ToUser: "user1|user2|user3",
    ToDept: "deptId1|deptId2|deptId3",
    PopWindow: &message.PopWindowItem{
        Url:      "url",//弹窗打开url
        Tip:      "tip",//提示内容
        Title:    "title",//窗口标题
        Width:    100,//弹窗宽度
        Height:   200,//弹窗宽度
        Duration: 6,//弹窗窗口停留时间。单位：秒，不设置或设置为0会取默认5秒, -1为永久
        Position: 1,//弹窗位置。 不设置或设置为0默认屏幕中央, 1 左上, 2 右上, 3 右下, 4 左下
        NoticeId: "12",//打开方式。1 浏览器, 2 窗口, 其他采用应用默认配置
        PopMode:  1,//通知id，用于防止重复弹窗
    },
}

//发送
yd.Message().Send(popWindowMessage)
```