package request

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"time"
)

type VideoFeedRequest struct {
	LatestTime int64  `json:"latest_time" form:"latest_time"`
	Token      string `json:"token" form:"token"`
}

func (v *VideoFeedRequest) VerifyFeed() {
	if v.LatestTime == 0 {
		v.LatestTime = time.Now().UnixMilli()
	}
}

type VideoPublishRequest struct {
	Data   *multipart.FileHeader `json:"data" form:"data"`
	Token  string                `json:"token" form:"token"`
	Title  string                `json:"title" form:"title"`
	UserId int64                 `json:"user_id" form:"user_id"`
}

func (v *VideoPublishRequest) VerifyFeed() (*bytes.Buffer, error) {
	if v.Data == nil {
		return nil, fmt.Errorf("找不到上传的文件呀~")
	}
	src, _ := v.Data.Open()
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, src)
	return buf, err
}

type PublishListRequest struct {
	Token  string `json:"token" form:"token"`
	UserId int64  `json:"user_id" form:"user_id"`
}

type FavoriteActionRequest struct {
	Token      string `json:"token" form:"token"`
	VideoId    int64  `json:"video_id" form:"video_id"`
	ActionType int32  `json:"action_type" form:"action_type"`
	UserId     int64  `json:"user_id" form:"user_id"`
}

type FavoriteListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type CommentActionRequest struct {
	UserId      int64  `json:"user_id" form:"user_id"`
	Token       string `json:"token" form:"token"`
	VideoId     int64  `json:"video_id" form:"video_id"`
	ActionType  int32  `json:"action_type" form:"action_type"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentId   int64  `json:"comment_id" form:"comment_id"`
}

type CommentListRequest struct {
	Token   string `json:"token" form:"token"`
	VideoId int64  `json:"video_id,omitempty" form:"video_id"`
	UserId  int64  `json:"user_id" form:"user_id"`
}
