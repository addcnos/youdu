[查看所有](./README.md)

## 群管理

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00015.html) 

### 创建群

```go
import "github.com/addcnos/youdu"

var name string = 'group-name'  // 群名称
groupId, err := yd.Group().Create(name)
```

### 删除群

```go
import "github.com/addcnos/youdu"

var groupId string = 'group-id' // 群ID
b, err := yd.Group().Delete(groupId)
```

### 修改群名称

```go
import "github.com/addcnos/youdu"

var id string = 'group-id' // 群ID
var name string = 'group-name' // 新的群名称
b, err := yd.Group().Update(id, name)
```

### 查看群信息

```go
import "github.com/addcnos/youdu"

var groupId string = 'group-id' // 群ID
group, err := yd.Group().Info(groupId)
```

### 群列表

指定用户ID获取群列表： 

```go
import "github.com/addcnos/youdu"

var ua string = 'a-user-id' // A用户ID
var ub string = 'b-user-id' // B用户ID
// 更多user id

groups, err := yd.Group().List(ua, ub)
```

不指定用户ID，获取所有群列表： 

```go
import "github.com/addcnos/youdu"

groups, err := yd.Group().List()
```

### 添加群成员

```go
import "github.com/addcnos/youdu"

var groupId string = 'group-id' // 群ID
var ua string = 'a-user-id' // A用户ID
var ub string = 'b-user-id' // B用户ID
// 更多user id

b, err := yd.Group().AddMember(groupId, ua, ub)
```

### 删除群成员

```go
import "github.com/addcnos/youdu"

var groupId string = 'group-id' // 群ID
var ua string = 'a-user-id' // A用户ID
var ub string = 'b-user-id' // B用户ID
// 更多user id

b, err := yd.Group().DelMember(groupId, ua, ub)
```

### 查询用户是否为群成员

```go
import "github.com/addcnos/youdu"

var groupId string = 'group-id' // 群ID
var userId string = 'user-id' // 用户ID

b, err := yd.Group().IsMember(groupId, userId)
```