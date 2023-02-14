package controller

import (
	"github.com/gin-gonic/gin"
	"simple-demo/common"
	"simple-demo/log"
	"simple-demo/response"
	"simple-demo/service"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	latest_time := c.Query("latest_time")
	uid := ""
	//如果token不为空的话,进行解析;
	if token != "" {
		err := common.CheckToken(token)
		if err != nil {
			response.Fail(c, CheckToken_Fail, nil)
			log.Error(err)
			return
		}
		uid, err = common.GetTokenUid(token)
		if err != nil {
			response.Fail(c, GetUid_Fail, nil)
			log.Error(err)
			return
		}
	}
	mess, err := service.GetFeedList(uid, latest_time)
	if err != nil {
		response.Fail(c, GetFeedList_Failed, err)
		log.Error(err)
		return
	}
	response.Success(c, GetFeedList_Success, mess)
}
