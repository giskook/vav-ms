# Create vehicle audio and vedio properties

Create vehicle audio and vedia properties

**URL** : `/vavms/api/v1/vehicle_property`

**Method** : POST

**Data constraints**

All fields must be sent

Filed audio format / audio sampling rate and vedio format is the index.The meaning is 1079 0x1003's content.

**Data example** 

```json
{ 
	"audio_format":"6",
	"audio_sampling_rate":"0"
	"vedio_format":"98",
}
```
## Success Response

**Condition** : vehicle property is set to redis correctly.

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

**Condition** : The fileds are over ranged

**Code** : `400 BAD REQUEST`

**Content** : 

```json
{
    "code":"40003",
    "desc":"音视频编码方式或音频采样率超出支持范围(40003)"
}
```

### OR

**Condition** : The audio or video formats are not supported

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50008",
    "desc":"暂不支持该格式，请联系对接(50008)"
}
```
### OR

**Condition** : access addr not add to redis

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50007",
    "desc":"设置车机音视频属性失败(50007)"
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
