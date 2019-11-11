# Show stream media

Show all stream medias

**URL** : `/vavms/api/v1/stream_media`

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
	"data":[]
}
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
	"stream_medias":[{"access_uuid":"vavms1","rtmp_application":"myapp","rtmp_ip_inner":"127.0.0.1", "rtmp_ip_outer":"222.222.218.52", "rtmp_port_inner":"8080", "rtmp_port_outter":"8023","http_location":"live","http_ip_outter":"222.222.218.51", "http_port_outter":"9002" }]
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
