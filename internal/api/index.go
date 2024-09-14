package api

import (
	"github.com/gin-gonic/gin"
)

func (m *ServiceServer) Index(c *gin.Context) {
	c.JSON(200, "see api documentation on http://localhost:3001/docs and try api on http://localhost:3001/explore/index.html")
}
