# curl
使用curl发送http请求

## post
* 普通
curl -d "username=sunlg&password=123456&email=sunlggggg@gmail.com" "http://127.0.0.1:8080/user"
* json 格式
curl -H "Content-Type:application/json" -X POST --data 'json' "http://www.baidu.com"



