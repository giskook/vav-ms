# stream media related 

## basic info

key : **vavms_stream_media** 

struct : **list** 

list entity is a json struct:

* **access_uuid**

* **domain_inner**

* **domain_outer**

### command 

`LRANGE vavms_stream_media 0 -1`

```json
 "{"access_uuid":"vavms1","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://192.168.2.122:8080/myapp"}"
2) "{"access_uuid":"vavms4","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://192.168.2.124:8080/myapp"}"
3) "{"access_uuid":"vavms3","domain_inner":"rtmp://127.0.0.1:8080/myapp","domain_outer":"rtmp://192.168.2.123:8080/myapp"}"

```
