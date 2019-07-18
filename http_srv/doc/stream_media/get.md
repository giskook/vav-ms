# Show stream media

Show all stream medias

**URL** : `/vavms/api/stream_media/`

**Method** : `GET`

**Data constraints** : `{}`

## Success Responses

**Condition** : No stream media

**Code** : `200 OK`

**Content** : 

```json
{ 
	"code":"0",
	"desc":"成功",
] }
```
### OR

**Condition** : Have stream media

**Code** : `200 OK`

**Content** :

```json
{
    "code":"0",
    "desc":"成功",
    "data":{
	"stream_medias":[{"access_uuid":"vavms1","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://222.222.218.52:8080/myapp" },
	{"access_uuid":"vavms1","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://222.222.218.53:8080/myapp" }]
}} 
```
## Error Responses

**Condition** : Can not read from redis

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50002",
    "desc":"读取流媒体服务器地址失败(50002)"
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
