# 个人资源服务
个人的文件之类可以上传到服务器,提供加密服务，提供cli客户端
## 功能完成情况
### 用户模块 
* [X] 注册 
* [X] json web token 
### 文件模块 
* [X] 文件上传
* [X] 二叉树和多叉树互转
* [ ] 缓存多叉树在内存中(后续可能上redis)
* [ ] 文件记入数据库
* [ ] 增加文件 
* [ ] 删除文件
* [ ] 增加文件夹
* [ ] 删除文件夹 
### 文本服务 
* [ ] 直接上传文本（考虑是否保存为文件，后续移植到数据库）
### 加密模块
* [ ] 文件加密
* [ ] 文本加密
### leveldb 
* [ ] 文件存入leveldb
### 本地服务 
* [ ] 本地文件管理，但是目录上传到服务端  
* [ ] 本地文件同步到服务端
* [ ] 本地文件加密
* [ ] 本地文本加密
## 进度管理
* [ ] 文件系统建立完毕 2018 8 19
## 更新记录
* 修改数据库表结构将二叉树存储转为普通多叉树存储 2018 8 20 
## go基础

* [go interface](docs/go%20interface.md)
* [go mysql](docs/go%20mysql使用.md)
* [go uc](docs/go%20uint%20case.md)
* [go http 实现restful](docs/go使用http实现restful.md)
* [go slice ](docs/go%20slice%20动态增删.md)
* [go struct](docs/go%20struct.md)
* [go json](docs/go使用json.md)

## curl
* [curl 使用](docs/curl%20usage.md)

## 树和二叉树的转换 

* [树和二叉树的转换](main/filetree/converter)

## jwt
使用json web token进行权限验证
并设置过期时间
* [实现](main/jwtutil)


## 客户端

提供给用户编译好的客户端，或者用户自行编译


## leveldb

将图片、文本存储在leveldb中

* [leveldb 使用](docs/leveldb.md)
