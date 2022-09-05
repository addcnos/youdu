[查看所有](../README.md)

## 素材管理

详细说明：[有度官方文档](https://youdu.im/doc/api/c01_00017.html) 

### 上传素材文件

```go
import "github.com/addcnos/youdu"

var fileType string = 'image' // 素材类型。image代表图片、file代表普通文件、voice代表语音、video代表视频
var filePath string = '/home/www/a.png' // 文件路径
mediaId, err := yd.Media().Upload(fileType, filePath)
```

### 下载素材文件

```go
import "github.com/addcnos/youdu"

var mediaId string = 'media-id' // 素材媒体文件ID
mediaGetResp, err := yd.Media().Get(mediaId)
```

### 查询素材文件信息

```go
import "github.com/addcnos/youdu"

var mediaId string = 'media-id' // 素材媒体文件ID
media, err := yd.Media().Search(mediaId)
```