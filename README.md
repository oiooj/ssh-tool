ssh-tool
========


Usage
=====

```shell
Usage of ssh-tool:
  -c="w": command you want to exec
  -h="192.168.0.1/24": ssh hosts, use CIDR
  -p="password": ssh user password
  -t=3: timout duration (s)
  -u="root": ssh user name
```

>./ssh-tool -u root -p yourpassword  -h 192.168.0/24 -t 2 -c "ps"

你可以自己编译，也可以下载上面的ssh-tool使用。
需要指出的是如果你自己编译直接引用 "github.com/flynn/go-crypto-ssh" 这个ssh库可能会有问题，因为他的库不支持自定义超时时间，为了加快执行的速度，同时让程序更加灵活，我加入了自定义超时时间，在他的库上做了些改进。



