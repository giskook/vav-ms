# Request audio and video stream

Request audio and video stream

**URL** : `/vavms/api/v1/stream/{sim}/{channel}`

**Method** : POST

**Data constraints**

**Data example** All fields must be sent

```json
{ 
	"type":"live/back",
}
```
## Success Response

### live stream

**Condition** : live stream already on

**Code** : `200 OK`

**Content example** :

```json 
{
    "code":"0",
    "desc":"成功",
    "data":{ 
		"url":"rtmp://222.222.218.52:8080/myapp/id_channel"
	}
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
