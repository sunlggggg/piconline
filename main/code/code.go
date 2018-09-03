package code

// 不是正常的get post put delete 请求
const (
	InvalidRequest        = "非法请求"
	UserExisted           = "用户已经存在"
	CreateSuccess         = "用户创建成功"
	FileRootCreateFail    = "目录初始化失败"
	FileRootCreateSuccess = "目录初始化成功"
	FileDirCreateSuccess  = "文件夹创建成功"
	FileDirCreateFail     = "文件夹创建失败"
	FileCreateSuccess     = "文件创建成功"
	FileCreateFail        = "文件创建失败"
	LoginErr              = "登陆出错"
	LoginSuccess          = "登陆成功"
	LoginFail             = "登陆失败"
	InValidToken          = "无效的Token"
	InValidArgs           = "前端参数错误"
)
