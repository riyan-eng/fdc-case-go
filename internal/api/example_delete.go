package api

import (
	"context"
	"server/infrastructure"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Delete
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id} [delete]
// @Security ApiKeyAuth
func (m *ServiceServer) ExampleDelete(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	err := m.dao.NewExampleRepository().Delete(&ctx, &id)
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]string{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_DELETE"))
}
