# 功能清单

## 企业应用

### 发送应用消息

- [x] 通用消息 `message.go@SendMessage`
- [x] 文本消息 `message.go@SendTextMessage`
- [x] 图片消息 `message.go@SendImageMessage`
- [x] 文件消息 `message.go@SendFileMessage`
- [x] 图文消息 `message.go@SendMpNewsMessage`
- [x] 隐式链接 `message.go@SendLinkMessage`
- [x] 外链消息 `message.go@SendExLinkMessage`
- [x] 系统消息 `message.go@SendSysMessage`
- [ ] 短信消息
- [ ] 邮件消息

### 设置角标

- [ ] 设置角标

### 发送弹窗消息

- [x] 发送弹窗消息 `message.go@SendPopWindowMessage`

## 单点登录

- [x] 单点登录 `identity.go@Identify`

## 会话管理

- [x] 创建会话 `session.go@CreateSession`
- [x] 获取会话 `session.go@GetSession`
- [x] 修改会话 `session.go@UpdateSession`

## 会话消息

- [x] 文本消息 `session.go@SendSessionMessage`
- [x] 图片消息 `session.go@SendSessionImageMessage`
- [x] 文件消息 `session.go@SendSessionFileMessage`
- [x] 语音消息 `session.go@SendSessionVoiceMessage`
- [x] 视频消息 `session.go@SendSessionVideoMessage`

## 自定义菜单

- [ ] 下载消息数据
- [ ] 第三方接口处理响应

## 企业通讯录

### 部门管理

- [x] 创建部门 `dept.go@CreateDept`
- [x] 更新部门 `dept.go@UpdateDept`
- [x] 删除部门 `dept.go@DeleteDept`
- [x] 获取部门列表 `dept.go@GetDeptList`
- [x] 获取部门 ID `dept.go@GetDeptAliasList` `dept.go@GetDeptId`

### 用户管理

- [x] 创建用户
- [x] 更新用户
- [x] 更新用户部门职务信息
- [x] 删除用户
- [x] 批量删除用户
- [x] 获取用户信息 `user.go@GetUser`
- [x] 获取部门用户详细信息 `user.go@GetDeptUserList`
- [x] 获取部门用户 `user.go@GetDeptUserSimpleList`
- [ ] 设置用户头像
- [ ] 获取用户头像
- [ ] 更新用户拓展属性字段
- [x] 查询用户激活状态 `user.go@GetUserEnableState`
- [x] 修改用户激活状态 `user.go@GetUserEnableStateUpdate`

### 第三方认证

- [ ] 设置认证信息

### 群管理

- [x] 创建群 `group.go@CreateGroup`
- [x] 删除群 `group.go@DeleteGroup`
- [x] 修改群名称 `group.go@UpdateGroup`
- [x] 查看群信息 `group.go@GetGroupInfo`
- [x] 群列表 `group.go@GetGroupList`
- [x] 添加群成员 `group.go@AddGroupMember`
- [x] 删除群成员 `group.go@DelGroupMember`
- [x] 查询用户是否为群成员 `group.go@IsGroupMember`

### 全量覆盖

- [ ] 发起全量覆盖
- [ ] 获取全量覆盖结果
- [ ] 全量覆盖完成通知

## 素材管理

- [ ] 上传素材文件
- [ ] 下载素材文件
- [ ] 查询素材文件信息

## 应用消息回调

- [x] 消息解密 `receive.go@Decrypt`
