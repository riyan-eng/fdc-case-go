package service

import (
	"context"
	"database/sql"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"
)

type ExampleService interface {
	Create(ctx *context.Context, ent *entity.ExampleCreate) *util.Error
	Patch(ctx *context.Context, ent *entity.ExamplePatch) *util.Error
	Put(ctx *context.Context, ent *entity.ExamplePut) *util.Error
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao *repository.DAO) ExampleService {
	return &exampleService{
		dao: *dao,
	}
}

func (m *exampleService) Create(ctx *context.Context, ent *entity.ExampleCreate) *util.Error {
	mdl := model.Example{
		Id:     *ent.Id,
		Name:   sql.NullString{String: *ent.Name, Valid: true},
		Detail: sql.NullString{String: *ent.Detail, Valid: util.NewIsValid().String(ent.Name)},
	}

	if err := m.dao.NewExampleRepository().Create(ctx, &mdl); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *exampleService) Patch(ctx *context.Context, ent *entity.ExamplePatch) *util.Error {
	mdl := model.Example{
		Id:     *ent.Id,
		Name:   sql.NullString{String: *ent.Name, Valid: util.NewIsValid().String(ent.Name)},
		Detail: sql.NullString{String: *ent.Detail, Valid: util.NewIsValid().String(ent.Detail)},
	}

	if err := m.dao.NewExampleRepository().Patch(ctx, &mdl); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *exampleService) Put(ctx *context.Context, ent *entity.ExamplePut) *util.Error {
	mdl := model.Example{
		Id:     *ent.Id,
		Name:   sql.NullString{String: *ent.Name, Valid: true},
		Detail: sql.NullString{String: *ent.Detail, Valid: util.NewIsValid().String(ent.Detail)},
	}

	if err := m.dao.NewExampleRepository().Put(ctx, &mdl); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}
