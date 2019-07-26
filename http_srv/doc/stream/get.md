# Request audio and video stream

Request audio and video stream

**URL** : `/vavms/api/v1/stream/{type}/{sim}/{channel}`

**Method** : GET

**Data constraints**

**Data example** All fields must be sent

"{}"

## Success Response

**Condition** : live stream on

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

**Condition** : stream is not upload 

**Code** : `504 GATEWAY TIMEOUT`

**Content** : 

```json
{
    "code":"50401",
    "desc":"音视频资源请求超时(50401)"
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
