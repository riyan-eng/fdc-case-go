package router

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.GET("", m.handler.ExampleList)
	subRoute.GET("/:id", m.handler.ExampleDetail)
	subRoute.POST("", m.handler.ExampleCreate)
	subRoute.DELETE("/:id", m.handler.ExampleDelete)
	subRoute.PUT("/:id", m.handler.ExamplePut)
	subRoute.PATCH("/:id", m.handler.ExamplePatch)
}
