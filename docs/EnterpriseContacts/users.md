## 用户管理

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00013.html) 

### 获取用户信息

```go
import "github.com/addcnos/youdu"

var userId string = 'user-id'   // 用户ID
user, err := youdu.NewUser().Get(userId)
```

### 获取部门用户详细信息

```go
import "github.com/addcnos/youdu"

var deptId int = 1  // 部门ID
users, err := youdu.NewUser().List(deptId)
```

### 获取部门用户

```go
import "github.com/addcnos/youdu"

var deptId int = 1  // 部门ID
simpleUsers, err := youdu.NewUser().SimpleList(deptId)
```

### 查询用户激活状态

```go
import "github.com/addcnos/youdu"

var userId string = 'user-id'   // 用户ID
state, err := youdu.NewUser().EnableState(userId)
```