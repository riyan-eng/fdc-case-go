package api

import (
	"context"
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.ExamplePut	true  "body"
// @Router		/example/{id} [put]
// @Security ApiKeyAuth
func (m *ServiceServer) ExamplePut(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	payload := new(dto.ExamplePut)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	if err := m.exampleService.Put(&ctx, &entity.ExamplePut{
		Id:     &id,
		Name:   &payload.Name,
		Detail: &payload.Detail,
	}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_UPDATE"), 200)
}
