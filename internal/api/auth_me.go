package api

import (
	"context"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Me
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/me [get]
// @Security 	BearerAuth
func (m *ServiceServer) AuthMe(c *gin.Context) {
	ctx := context.Background()
	user := util.CurrentUser(c)

	data, err := m.authService.Me(&ctx, &entity.AuthMe{UserId: &user.UserId})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
