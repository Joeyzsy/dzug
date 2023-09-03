package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendUser struct {
	Msg           string `json:"message"`
	MsgType       int64  `json:"msgType"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注

	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
}

type GetUserInfoListResp struct {
	Response
	UserInfos []*User `json:"user_list"`
}

type GetFriendInfoListResp struct {
	Response
	FriendInfos []*FriendUser `json:"user_list"`
}

type PostRelationResp struct {
	Response
}

func UserInfoListRespSuccess(c *gin.Context, list []*User) {
	c.JSON(http.StatusOK, GetUserInfoListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		UserInfos: list,
	})
}

func FriendInfoListRespSuccess(c *gin.Context, list []*FriendUser) {
	for _, v := range list {
		fmt.Printf("*********************%++v", v)
	}

	c.JSON(http.StatusOK, GetFriendInfoListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		FriendInfos: list,
	})
}

func PostRelationRespSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, PostMessageResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
	})
}
