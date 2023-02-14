package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"simple-demo/common"
	"simple-demo/config"
	"simple-demo/log"
	"simple-demo/response"
	"simple-demo/service"
	"simple-demo/util"
)

// Publish check token then save upload file to public directory
func PublishAction(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	err := common.CheckToken(token)
	if err != nil {
		response.Fail(c, CheckToken_Fail, nil)
		log.Error(err)
		return
	}
	uid, err := common.GetTokenUid(token)
	if err != nil {
		response.Fail(c, GetUid_Fail, nil)
		log.Error(err)
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		response.Fail(c, Publish_Fail, err)
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s", util.RandomString(), filename)
	savePath := config.Config.Path.Videofile
	saveFile := filepath.Join(savePath, finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		response.Fail(c, SaveVideo_Fail, err)
		log.Error(err)
		return
	}
	mess, err := service.PublishAction(uid, title, saveFile)
	if err != nil {
		response.Fail(c, SaveVideo_Fail, err)
		log.Error(err)
		return
	}
	response.Success(c, Publish_Success, mess)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	err := common.CheckToken(token)
	if err != nil {
		response.Fail(c, CheckToken_Fail, nil)
		return
	}
	mess, err := service.PublishList(uid)
	if err == nil {
		response.Success(c, PublishList_Success, mess)
	} else {
		response.Fail(c, PublishList_Failed, mess)
		return
	}
}
