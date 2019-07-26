# stream related 

## basic info

key : **id_channel_play_type** `13731143001_2_live/13731143001_2_back`

struct : **hash** 

hash entity:

* **data_type**

* **uuid**

* **ttl**

* **url**

### command 

`HGETALL 13731143001_2_live`

```
"data_type":"0"
"uuid":"vavms1"
"ttl":"600"
"url":"rtmp://222.222.218.52:8080/myapp/13731143001_2"
```
### life circle

When `POST vavms/api/v1/stream` set the `key` and `data_type`

When stream uploaded use `data_type` to create pipe and exec ffmpeg bin, set 

the `uuid`, use `ttl` to expire the key, set `url` for the http

When `PUT vavms/api/v1/stream` update the key's ttl 

## status info

status info is used to keep the process atomtic

key : **id_channel_status** `13731143001_2_status` must be "live" or "back"

struct: **key/value**

entity:

* **id_channel_status**

## command

`GET 13731143001_2_status`

```
"live"
```

### life circle

When `POST vavms/api/v1/stream` set the status

When stream uploaded del the key
