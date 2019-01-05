# android-agent
android端http服务器，通过wifi接收http请求
后期可以加入adb shell命令获取手机信息

## 编译(android手机)
```
go get github.com/sevlyar/go-daemon
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build

```

## 使用
```
# adb连接手机
adb devices
adb push  android-agent  /data/local/tmp
adb shell chmod 755 /data/local/tmp/android-agent
adb shell /data/local/tmp/android-agent

```

## 后续
- 查端口号  
netstat -tunlp  

- 查进程  
adb shell  
ps | grep android  

- 杀进程  
kill -s 9 1827  
