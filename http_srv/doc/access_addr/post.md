# Create access addr

Create access addr for vavms

**URL** : `/vavms/api/v1/access_addr`

**Method** : POST

**Data constraints**

All fields must be sent

**Data example** 

```json
{ 
	"IP":"222.222.218.52",
	"Port":"8875"
}
```
**curl**
```
curl -H "Content-Type:application/json" -d '{ "IP":"222.222.218.51", "Port":"8875"}' http://127.0.0.1:9001/vavms/api/v1/access_addr
```
## Success Response

**Condition** : access addr is set to redis correctly.

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

**Condition** : access addr not add to redis

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50005",
    "desc":"设置车机连接地址失败(50005)"
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
