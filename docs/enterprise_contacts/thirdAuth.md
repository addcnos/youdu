[查看所有](./README.md)

## 第三方认证

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00014.html) 

### 设置认证信息

```go
import "github.com/addcnos/youdu"

var userId string = 'user-id'   // 用户ID
var password string = 'password'    // 密码
b, err := yd.Auth().SetAuth(userId, password)
```