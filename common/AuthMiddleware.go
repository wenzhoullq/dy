package common

// 生成token
func GenToken(username string, password string) (string, error) {
	token := ""
	return token, nil
}

// 账号密码参数检验
func Check_AccountParam(username string, password string) bool {
	//防sql注入
	//账户长度不得大于32位,密码长度不得大于32位
	if len(username) > 32 || len(password) > 32 {
		return false
	}
	return true
}
