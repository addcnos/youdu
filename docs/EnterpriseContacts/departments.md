[查看所有](https://github.com/addcnos/youdu#%E8%AF%A6%E7%BB%86%E6%96%87%E6%A1%A3)

## 部门管理

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00012.html) 

### 获取部门列表

```go
import "github.com/addcnos/youdu"

var deptId int = 1 // 部门ID
depts, err := youdu.NewDept().GetList(deptId)
```