package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthRegister	true  "body"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/register [post]
func (m *ServiceServer) AuthRegister(c *gin.Context) {
	ctx := context.Background()
	payload := new(dto.AuthRegister)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	id := uuid.NewString()
	roleCode := "STAF"
	err := m.authService.Register(&ctx, &entity.AuthRegister{
		UserId:   &id,
		UserName: &payload.Username,
		Password: &payload.Password,
		Email:    &payload.Email,
		RoleCode: &roleCode,
	})

	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id":       id,
		"username": payload.Username,
		"email":    payload.Email,
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
