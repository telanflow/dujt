# 毒鸡汤 - Go
壮士可要来一碗！

## 安装
[Download](https://github.com/telanflow/dujt/releases)
下载可执行文件到 dujt/web/ 目录执行即可。

## 特性
- 超快响应、性能怪兽
- 部署简单
- 无数据库依赖（默认：毒鸡汤文件放在 web/statement.txt 中，根据换行识别一条语句）
- 自动重载：程序会监听 “毒鸡汤文件（statement.txt）” 变更，然后重新加载，无需重启程序

## Usage
```
➜  ./darwin_amd64 -h

Usage of ./darwin_amd64:
  -addr string
    	服务地址 (default "0.0.0.0")
  -port string
    	服务端口 (default "80")
  -target string
    	毒鸡汤文件路径 (default "statement.txt")
``` 

## 对于维护：
1. 直接修改 web/statement.txt 文件，添加一条语句即可，无需重启程序！
2. 需要结束程序  直接 kill 即可
3. Go程序后台执行，由于我没有写此功能，so 可以借助第三方实现：

> 守护进程管理工具：supervisor

> Linux窗口管理工具：screen

> linux_amd64 -addr 0.0.0.0 -port 80 > /dev/null 2>&1

## 变量含义 (web/index.html)
    url: http://localhost
    host: localhost
    line: 毒鸡汤语句

## 预览图
![screenshot](screenshot_exec.jpg)
![screenshot](screenshot.png)

## 感谢
[原版毒鸡汤 PHP](https://github.com/egotong/nows)
