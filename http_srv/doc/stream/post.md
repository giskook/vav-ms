# Set stream class [live] or [back]

Set stream class [live] or [back]

**URL** : `/vavms/api/v1/stream/{type}/{sim}/{channel}`

**Method** : POST

**Data constraints**

json format sub keys `data_type` and `time_out`

All fields must be sent

`data_type` accept values

0 stands for audio and video

1 stands for video

2 stands for two way intercom

3 stands for listen

4 stands for broadcast

5 stans for tranparent transmission

`ttl` units second

`number` 车牌号码

`priority` 如果是监管传入>=100的值,如非传入0-100

`token` 监管请求视频需要设置

**Data example** 
```json
{
	"data_type":"0",
	"ttl":"1000",
	"priority":"101",
	"number":"冀AD8V52",
	"token":"ddxdbgda1"
}
```
**curl**
```
curl -H "Content-Type:application/json" -d '{ "data_type":"0", "ttl":"1000", "priority":"101", "number":"冀AD8V52", "token":"ddxdbgda1" }' http://127.0.0.1:9001/vavms/api/v1/stream/2/17190604001/1
```

## Success Response

**Condition** : 

**Code** : `200 OK`

**Content example** :

```json 
{
    "code":"0",
    "desc":"成功",
    "data":{
	"ip":"222.222.218.52",
	"port":"8876"
	}
}
```
### OR

**Condition** : url already exists only for the live

**Code** : `200 Accepted`

**Content example** :

```json 
{
    "code":"20000",
    "desc":"url已存在",
    "data":"rtmp://222.222.218.52:8888/myapp/13731143001_1_live"
}
```
### OR
**Condition** : url already exists only for the back 

**Code** : `200 Accepted`

**Content example** :

```json 
{
    "code":"20001",
    "desc":"另外一个用户正在观看，请等待",
}
```
### OR
**Condition** : vehicle's audio format and video format are not configured

**Code** : `202 Accepted`

**Content example** :

```json 
{
    "code":"20200",
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
    "code":"20202",
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

**Condition** : The fields are not 

**Code** : `400 BAD REQUEST`

**Content** : 

```json
{
    "code":"40002",
    "desc":"参数值非法(40002)"
}
```
### OR 

**Condition** : another stream media is show

**Code** : `409 Conflict`

**Content** : 

```json
{
    "code":"40900",
    "desc":"另外一个用户正在请求，请稍等(40900)"
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
