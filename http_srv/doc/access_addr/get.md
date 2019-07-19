# Show access addr

Show access addr(LVS addr) for vavms

**URL** : `/vavms/api/v1/access_addr`

**Method** : `GET`

**Data constraints** : `{}`

## Success Responses

**Condition** : No access addr

**Code** : `200 OK`

**Content** : 

```json
{ 
	"code":"0",
	"desc":"成功",
	"data":{
		"ip":"",
		"port":""
	}
}
```
### OR

**Condition** : Have access addr

**Code** : `200 OK`

**Content** :

```json
{
    "code":"0",
    "desc":"成功",
    "data":{
	"ip":"222.222.218.52",
	"port":"8875"
    }
} 
```
## Error Responses

**Condition** : Can not read from redis

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50006",
    "desc":"得到车机连接地址失败(50002)"
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
