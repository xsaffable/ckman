package user

import (
	"github.com/gin-gonic/gin"
	"github.com/housepower/ckman/common"
)

// 定义角色常量
const (
	ADMIN    = "admin"
	ORDINARY = "ordinary"
)

// Role UserRole 定义角色对应的标识
var Role = map[string]int32{
	ADMIN:    0,
	ORDINARY: 1,
}

// 获取用户角色标识
func getUserRole(c *gin.Context, role string) int32 {
	roleIdentify, ok := Role[role]
	if ok {
		return roleIdentify
	}
	return Role[ORDINARY]
}

// 根据用户获取角色
func getUserRoleByUsername(c *gin.Context, username string) int32 {
	if common.DefaultUserName == username {
		return getUserRole(c, ADMIN)
	}
	return getUserRole(c, ORDINARY)
}
