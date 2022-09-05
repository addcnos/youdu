[查看所有](../README.md)

## 单点登录

参考 [有度即时通讯服务端 API 文档 -- 单点登录](https://youdu.im/doc/api/c01_00007.html)

### 登录接口


```go
token, err := yd.AccessToken() //获取token
if err != nil {
    return err
}
response :=yd.Auth().Identify(token)

```
返回值类型

```go

type IdentifyResp struct {
	Buin   int `json:"buin"`
	Status struct {
		Code      int    `json:"code"`//状态码
		Message   string `json:"message"`//状态码相关的信息描述
		CreatedAt string `json:"createdAt"`//结果返回时间
	} `json:"status"`
	UserInfo struct {
		Gid        int    `json:"gid"`
		Account    string `json:"account"`//帐号
		ChsName    string `json:"chsName"`//中文名
		EngName    string `json:"engName"`//英文名
		Gender     int    `json:"gender"`//性别。0：男；1：女
		OrgId      int    `json:"orgId"`
		Mobile     string `json:"mobile"`//手机号码
		Phone      string `json:"phone"`//电话号码
		Email      string `json:"email"`//邮箱
		CustomAttr string `json:"customAttr"`
	} `json:"userInfo"`
}
```

返回值数据

```go
{
  "status": {
      "code": 0,
      "message": "Action completed successful",
      "createdAt": "2017-01-01 10:00:00"
  },
  "userInfo": {
      "account": "$account",
      "chsName":"$chs_name",
      "gender":1,
      "mobile":"$mobile",
      "phone":"$phone",
      "email":"$email"
  }
}

```