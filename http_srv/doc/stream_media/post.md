# Create stream media

Create stream media for vavms

**URL** : `/vavms/api/stream_media/`

**Method** : POST

**Data constraints**

Provide stream media 

```json
{ 
	"stream_medias":[{"access_uuid":"[vavms1]","domain_inner":"[rtmp inner addr]","domain_outer":"[rtmp outer addr]" },
	"stream_medias":[{"access_uuid":"[vavms2]","domain_inner":"[rtmp inner addr]","domain_outer":"[rtmp outer addr]" },
] }
```

**Data example** All fields must be sent

```json
{ 
	"stream_medias":[{"access_uuid":"vavms1","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://222.222.218.52:8080/myapp" },
	{"access_uuid":"vavms1","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://222.222.218.53:8080/myapp" }]
 }
```
## Success Response

**Condition** : stream media is set to redis correctly.

**Code** : `201 Created`

**Content example**

```json
{
    "code":"0",
    "desc":"成功"
}
```

## Error Response

**Condition** : The paramter is not decode correctly

**Code** : `400 BAD REQUEST`

**Content** : 

```json
{
    "code":"40000",
    "desc":"参数解析出错(40000)"
}
```

### OR

**Condition** : The fields are missed

**Code** : `400 BAD REQUEST`

**Content** : 

```json
{
    "code":"40001",
    "desc":"缺少参数(40001)"
}
```

### OR

**Condition** : stream media not add to redis

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50001",
    "desc":"添加流媒体服务器地址失败(50001)"
}
```
### OR

**Condition** : panic

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50000",
    "desc":"内部错误(50000)"
}
```

