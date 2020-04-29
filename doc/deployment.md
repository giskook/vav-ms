# 部署手册

## 软件集成

| 被集成的单元 |   环境要求 | 集成测试要求 | 完成准则 |
| -------- | ---------- | ---- | ---- |
| FFmpeg | CentOS7 | 启动FFmpeg | 有输出 |
| Nginx(nginx-http-flv-module| CentOS7 |启动测试脚本，尝试发布视频，并使用VLC观看视频|VLC正常输出|
| Redis| CentOS7 |redis重启后数据不丢失|设置redis里面的值，重启redis，数据不丢失|
| 音视频接入服务 | CentOS7(内核版本>2.6.35) | 启动程序正常，挂载车机到指定端口|能够看到车机实时音视频|

[FFmpeg安装](https://ffmpeg.org/download.html)
[nginx-http-flv-module安装](https://github.com/winshining/nginx-http-flv-module)
[redis安装](https://redis.io/)
