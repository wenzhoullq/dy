package controller

const (
	Check_AccountParam_Fail string = "账号密码格式错误"
	UserRegister_Fail       string = "账号注册失败,请稍后再试"
	UserRegister_Success    string = "账号注册成功,即将跳转新的页面"
	UserInfo_Success        string = "获取用户成功"
	UserInfo_Fail           string = "获取用户失败"
	UserLogin_Success       string = "账号登陆成功,即将跳转新的页面"
	UserLogin_Fail          string = "账号不存在或者密码错误"
	CheckToken_Fail         string = "验证token失败"
	GetUid_Fail             string = "解析Token获得Uid失败"
	PublishList_Success     string = "获得发布列表成功"
	PublishList_Failed      string = "获得发布列表失败"
	Publish_Fail            string = "发布视频失败"
	Publish_Success         string = "发布视频成功"
	SaveVideo_Fail          string = "保存视频失败"
	GetFeedList_Success     string = "获得视频流成功"
	GetFeedList_Failed      string = "获得视频流视频"
)
