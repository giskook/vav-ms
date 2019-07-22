# Request audio and video stream

Request audio and video stream

**URL** : `/vavms/api/v1/stream/{sim}/{channel}`

**Method** : POST

**Data constraints**

**Data example** All fields must be sent

```json
{ 
	"type":"live/back",
	"priority":0
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
### OR

**Condition** : live stream not upload

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20200",
    "desc":"请下发1078音视频指令",
}
```
### OR

**Condition** : 1078 alread sent please wait. 

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20201",
    "desc":"1078音视频指令已下发,请等待",
}
```
### OR

**Condition** : vehicle's audio format and video format are not configured

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20202",
    "desc":"车机音视频格式未设置",
}
```
### OR

**Condition** : access addr are not set

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20203",
    "desc":"车机接入地址未设置",
}
```
### OR

**Condition** : stream media are not set 

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20204",
    "desc":"流媒体地址未设置",
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

**Condition** : panic

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50000",
    "desc":"内部错误(50000)"
}
```

