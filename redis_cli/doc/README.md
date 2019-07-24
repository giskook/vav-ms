# Vavms redis store struct

## stream media related 

The key is **vavms_stream_media** use **list** struct

the list entity is a json struct.the json include 
**access_uuid**
**domain_inner**
**domain_outer**

## access addr related

* [Set access addr](access_addr/post.md):`POST /vavms/api/v1/access_addr`
* [Get access addr](access_addr/get.md):`GET /vavms/api/v1/access_addr`

## vehicle audio and video properties

* [Set access addr](vehicle_property/post.md):`POST /vavms/api/v1/vehicle_property`

## request video and audio stream

* [Set stream class](stream/post.md):`POST /vavms/api/v1/stream/{sim}/{channel}`
* [Request audio and video stream](stream/get.md):`GET /vavms/api/v1/stream/{sim}/{channel}`

