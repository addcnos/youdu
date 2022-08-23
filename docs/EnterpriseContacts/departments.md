## 部门管理

### 获取部门列表

```go
import "github.com/addcnos/youdu"

var depId int = 1
depts, err := youdu.NewDept().GetList(depId)
```