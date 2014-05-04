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



