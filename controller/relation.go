package controller

import (
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {

}
