package service

const (
	InitStatusCode        int32  = 0
	InitStatusMsg         string = "success"
	InitUserId            int64  = 0
	InitToken             string = ""
	FailStatusCode        int32  = -1
	UserNameExitMsg       string = "该用户名已存在,请修改用户名后再注册"
	UserNameNotExitMsg    string = "该用户名不存在,请检查账号密码"
	UserLoginFailed       string = "账号密码错误,登陆失败"
	UserLoginSuccess      string = "登陆成功，即将为你跳转新的页面"
	FailRegisterUser      string = "注册失败,请重试"
	FailInsertUser        string = "插入用户失败"
	FailGetUid            string = "获得UID失败"
	FailGenerateToken     string = "创建Token失败"
	FailGenerateCache     string = "插入缓存失败"
	FailGetUserinfo       string = "获取用户信息失败"
	KeyUserinfo           string = "userInfo_uid_"
	GetPublishListSuccess        = "获得发布列表成功"
	GetPublishListFailed         = "获得发布列表失败"
	KeyPublishList               = "uid_publish_"
	UserPublishSuccess           = "视频发布成功"
	PublishFail                  = "视频发布失败"
	CreateCoverUrlFail           = "生成视频封面失败"
	FailGetFeedList              = "获得视频流失败"
)
