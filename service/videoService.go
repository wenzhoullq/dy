package service

import (
	"os/exec"
	"path/filepath"
	"simple-demo/common"
	"simple-demo/config"
	"simple-demo/log"
	mes "simple-demo/proto/pkg"
	"simple-demo/repository"
	"strings"
)

func changePublistList(videolist []repository.Video) []*mes.Video {
	videoList := changeVideosForm(videolist)
	return videoList
}

func PublishList(uid string) (*mes.DouyinPublishListResponse, error) {
	mes := &mes.DouyinPublishListResponse{
		StatusCode: InitStatusCode,
		StatusMsg:  GetPublishListSuccess,
	}
	videoList, err := repository.PublishList(uid)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = GetPublishListFailed
		log.Error(err)
		return mes, err
	}
	mes.VideoList = changePublistList(videoList)
	m := make(map[string]string)
	return mes, nil
}
func PublishAction(uid string, title string, videofilePath string) (*mes.DouyinPublishActionResponse, error) {
	mes := &mes.DouyinPublishActionResponse{
		StatusCode: InitStatusCode,
		StatusMsg:  UserPublishSuccess,
	}
	coverUrl, err := CreateCoverUrl(videofilePath, uid)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = CreateCoverUrlFail
		return mes, err
	}
	videoUrl, err := CreateVideoUrl(videofilePath, uid)
	err = repository.InsertVideo(uid, title, coverUrl, videoUrl)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = PublishFail
		log.Error("insert video fail ", err)
		return mes, err
	}
	return mes, nil
}

func CreateVideoUrl(videofilePath string, uid string) (string, error) {
	videoUrl := ""
	//视频上传到minio
	client := common.GetMinio()
	videoUrl, err := client.UploadFile("video", videofilePath, uid)
	return videoUrl, err
}

func CreateCoverUrl(videofilePath string, uid string) (string, error) {
	coverUrl := ""
	imageFile, err := GetImageFile(videofilePath)
	if err != nil {
		log.Error("create_coverUrl failed:", err)
		return coverUrl, err
	}
	//封面上传到minio
	client := common.GetMinio()
	coverUrl, err = client.UploadFile("pic", imageFile, uid)
	return coverUrl, err
}

func GetImageFile(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "/")
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	picName := string(b[:len(b)-3]) + "png"
	picpath := config.GetConfig().Path.Picfile
	picfile := filepath.Join(picpath, picName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picfile)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	log.Debugf(picName)
	return picfile, nil
}
func GetFeedList(uid, latest_time string) (*mes.DouyinFeedResponse, error) {
	mes := &mes.DouyinFeedResponse{
		StatusCode: InitStatusCode,
		StatusMsg:  UserPublishSuccess,
	}
	videos, err := repository.GetFeedList(latest_time)
	if err != nil {
		mes.StatusCode = FailStatusCode
		mes.StatusMsg = FailGetFeedList
	}
	if uid != "" {

	}
	mes.VideoList = changeVideosForm(videos)
	return mes, nil
}
func changeVideosForm(video_proto []repository.Video) []*mes.Video {
	videos := make([]*mes.Video, len(video_proto))
	for i, video := range video_proto {
		Author := &mes.User{
			Id:            video.Author.Id,
			Name:          video.Author.Name,
			FollowCount:   video.Author.Follow_count,
			FollowerCount: video.Author.Follower_count,
			IsFollow:      video.Author.Is_follow,
		}
		v := mes.Video{
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Author:        Author,
			Title:         video.Title,
		}
		videos[i] = &v
	}
	return videos
}
