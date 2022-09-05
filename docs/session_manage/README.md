[查看所有](../README.md)

## 会话管理

参考 [有度即时通讯服务端 API 文档 -- 会话管理](https://youdu.im/doc/api/c01_00008.html)


### 创建会话

```go
title := "session title"                    // 会话标题。最多允许64个字符
members := []string{"mem1", "mem2", "mem3"} // 会话成员账号列表。包括创建者，多人会话的成员数必须在3人及以上
response, _ := yd.Session().Create(title, members)

```
##### 返回值类型

```go
type Session struct {
	SessionId  string   `json:"sessionId"`//会话id
	Type       string   `json:"type"`//会话类型。仅支持多人会话(multi)
	Owner      string   `json:"owner"`//会话创建者账号
	Title      string   `json:"title"`//会话标题
	Version    int      `json:"version"`//会话版本号，64位整型
	Member     []string `json:"member"`//会话成员账号列表。包括创建者
	LastMsgId  int      `json:"lastMsgId"`
	ActiveTime int      `json:"activeTime"`
}

```

##### 返回值数据

```go
{
    "sessionId": "$session_id",
    "title": "$session_title",
    "owner": "$creator",
    "version": 1,
    "type": "multi",
    "member": [
        "$mem1",
        "$mem2",
        "$mem3"
    ]
}

```

### 获取会话
```go
sessionId := "sessionId" //会话id
response, _ := yd.Session().Get(sessionId) //返回值 为 上文提过的 Session 类型
```
返回值数据见上文 [返回值类型](#返回值类型)


### 修改会话
```go
sessionId := "sessionId" ///会话id
opUser := "user1" //操作者账号
title := "title"//标题
addMembers := []string{"mem3", "mem4"} //新增会话成员账号列表
delMembers := []string{"mem1", "mem2"} //删除会话成员账号列表

response,_ := yd.Session().Update(sessionId, opUser, title, addMembers, delMembers)//返回值 为 上文提过的 Session 类型
```
返回值数据见上文 [返回值类型](#返回值类型)