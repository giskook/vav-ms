# Modify stream media

Modify the specific stream media

**URL** : `vavms/api/v1/stream_media/{index}`

**URL Parameters** : `index=[string]` where `index` is the index of the stream media in redis

**Method** : `PUT`

**Data** : 

```json
{
	"access_uuid":"vavms1",
	"domain_inner":"rtmp://127.0.0.1:8080/myapp",
	"domain_outer":"rtmp://222.222.218.52:8080/myapp"
}
```

## Success Response

**Condition** : If the index exists.

**Code** : `200 OK`

**Content** : 

```json
{
    "code":"0",
    "desc":"成功"
}
```
## Error Responses

**Condition** : update error

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50004",
    "desc":"更新流媒体服务器地址失败(50003)"
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
