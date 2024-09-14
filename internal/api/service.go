package api

import (
	"server/internal/repository"
	"server/internal/service"
)

type ServiceServer struct {
	dao              repository.DAO
	exampleService   service.ExampleService
	authService      service.AuthService
	perangkatService service.PerangkatService
	objectService    service.ObjectService
}

func NewService(
	dao repository.DAO,
	exampleService service.ExampleService,
	authService service.AuthService,
	perangkatService service.PerangkatService,
	objectService service.ObjectService,
) *ServiceServer {
	return &ServiceServer{
		dao:              dao,
		exampleService:   exampleService,
		authService:      authService,
		perangkatService: perangkatService,
		objectService:    objectService,
	}
}
