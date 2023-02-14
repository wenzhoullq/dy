package repository

import (
	"simple-demo/common"
	"simple-demo/util"
	"strconv"
)

type Video struct {
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	IsFavorite    bool   `gorm:"column:is_favorite;"`
	Author        User   `gorm:"column:user;"`
}

func InsertVideo(uid string, title string, coverUrl, playUrl string) error {
	uid_int, _ := strconv.ParseInt(uid, 10, 64)
	video := Video{
		AuthorId:      uid_int,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   util.GetCurrentTime(),
		Title:         title,
		IsFavorite:    false,
	}
	db := common.GetDB()
	err := db.Table("videos").Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
func PublishList(uid string) ([]Video, error) {
	var videos []Video
	uid_int, _ := strconv.ParseInt(uid, 10, 64)
	author, err := GetUser(uid_int)
	if err != nil {
		return videos, err
	}
	db := common.GetDB()
	err = db.Table("videos").Where("author_id = ?", uid).Order("video_id DESC").Find(&videos).Error
	if err != nil {
		return videos, err
	}
	for _, v := range videos {
		v.Author = author
	}
	return videos, nil
}
func GetFeedList(latest_time string) ([]Video, error) {
	var videos []Video
	db := common.GetDB()
	err := db.Table("videos").Order("publish_time DESC").Find(&videos).Error
	if err == nil {
		return videos, err
	}
	return videos, nil
}
