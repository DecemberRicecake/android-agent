# android-agent
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7

cd F:\android-agent
adb push  android-agent  /data/local/tmp
adb shell chmod 755 /data/local/tmp/android-agent
adb shell /data/local/tmp/android-agent


安装android SDK/ndk，将CC=android SDK中自带的gcc编译器。


查端口号  netstat -tunlp
查进程：
adb shell
ps | grep android
杀进程：
kill -s 9 1827