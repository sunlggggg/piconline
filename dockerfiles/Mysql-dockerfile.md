# 创建Dockerfile

首先，创建目录mysql-piconline,用于存放后面的相关东西。
```
$ mkdir -p data logs conf
```

data目录将映射为mysql容器配置的数据文件存放路径

logs目录将映射为mysql容器的日志目录

conf目录里的配置文件将映射为mysql容器的配置文件

进入创建的mysql目录，创建Dockerfile

# 创建镜像
```
$ docker build -t  -piconline .
```

# 查看镜像
```
$ docker images |grep mysql
```
# 运行容器
