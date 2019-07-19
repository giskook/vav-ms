# Delete stream media

Delete the specific stream media

**URL** : `vavms/api/v1/stream_media/{index}`

**URL Parameters** : `index=[string]` where `index` is the index of the stream media in redis

**Method** : `DELETE`

**Data** : `{}`

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

**Condition** : del error

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : 

```json
{
    "code":"50003",
    "desc":"删除流媒体服务器地址失败(50003)"
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
