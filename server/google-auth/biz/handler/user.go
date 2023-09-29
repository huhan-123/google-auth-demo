package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google-auth/biz/middleware"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	userId, exists := c.Get(middleware.AuthMiddleware.IdentityKey)
	if !exists {
		c.JSON(consts.StatusUnauthorized, utils.H{"message": "please login first"})
		return
	}

	// get user_info from database
	user := middleware.GetUserById(userId.(int64))

	c.JSON(consts.StatusOK, utils.H{
		"code": consts.StatusOK,
		"msg":  "",
		"data": user,
	})
}
