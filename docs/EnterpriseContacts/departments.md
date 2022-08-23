## 部门管理

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00012.html) 

### 获取部门列表

```go
import "github.com/addcnos/youdu"

var depId int = 1
depts, err := youdu.NewDept().GetList(depId)
```