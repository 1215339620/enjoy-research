package model

import "gorm.io/gorm"

type TreeHole struct {
	gorm.Model
	UserID      uint              `json:"userID"` //用户ID
	Content     string            //树洞内容
	LikeNum     uint              //点赞数
	Replays     []TreeHoleComment //回复的评论
	IsAnonymity uint              //是否匿名
}
type TreeHoleComment struct {
	gorm.Model
	CommentReplays    []TreeHoleComment
	UserID            uint   `json:"userID"` //用户ID
	TreeHoleCommentID uint   //评论父级ID  0
	Content           string //内容
	TreeHoleID        uint   //树洞ID
	IsAnonymityReplay uint   //是否匿名回复
}
