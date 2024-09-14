package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/infrastructure"
	"server/internal/datastruct"
	"server/internal/model"
	"server/util"
	"strings"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type ExampleRepository interface {
	List(ctx *context.Context, search *string, limit, offset *int) (*[]datastruct.ExampleList, *int, *util.Error)
	Create(ctx *context.Context, mdl *model.Example) *util.Error
	Patch(ctx *context.Context, mdl *model.Example) *util.Error
	Put(ctx *context.Context, mdl *model.Example) *util.Error
	Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error)
	Delete(ctx *context.Context, id *string) *util.Error
}

type exampleRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *exampleRepository) Create(ctx *context.Context, mdl *model.Example) *util.Error {
	sqlRslt, errT := m.sqlxDB.NamedExecContext(*ctx, `
	insert into example (uuid, name, detail) values (:id, :name, :detail)
	`, mdl)
	if errT != nil {
		if strings.Contains(errT.Error(), `duplicate key value violates unique constraint`) {
			return &util.Error{
				Errors:     "duplicate",
				Message:    "name has been used.",
				StatusCode: 409,
			}
		}
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no rows",
			Message:    infrastructure.Localize("FAILED_CREATE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}

func (m *exampleRepository) Patch(ctx *context.Context, mdl *model.Example) *util.Error {
	sqlRslt, errT := m.sqlxDB.NamedExecContext(*ctx, `
	update example set name=coalesce(:name, name), detail=coalesce(:detail, detail), updated_at=now() where uuid=:id
	`, mdl)
	if errT != nil {
		if strings.Contains(errT.Error(), `duplicate key value violates unique constraint`) {
			return &util.Error{
				Errors:     "duplicate",
				Message:    "name has been used.",
				StatusCode: 409,
			}
		}
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no rows",
			Message:    infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}

func (m *exampleRepository) Put(ctx *context.Context, mdl *model.Example) *util.Error {
	sqlRslt, errT := m.sqlxDB.NamedExecContext(*ctx, `
	update example set name=:name, detail=:detail, updated_at=now() where uuid=:id
	`, mdl)
	if errT != nil {
		if strings.Contains(errT.Error(), `duplicate key value violates unique constraint`) {
			return &util.Error{
				Errors:     "duplicate",
				Message:    "name has been used.",
				StatusCode: 409,
			}
		}
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no rows",
			Message:    infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}

func (m *exampleRepository) List(ctx *context.Context, search *string, limit, offset *int) (*[]datastruct.ExampleList, *int, *util.Error) {
	data := make([]datastruct.ExampleList, 0)
	countRow := new(int)

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail, count(e.uuid) over() as total_rows
	from example e 
	where lower(e.name) like lower('%%%v%%') order by e.created_at %v limit %v offset %v
	`, *search, "desc", *limit, *offset)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return &data, countRow, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Rows(&data, sqlRows); err != nil {
		return &data, countRow, &util.Error{
			Errors: err.Error(),
		}
	}

	for _, d := range data {
		countRow = &d.TotalRows
		break
	}

	return &data, countRow, &util.Error{}
}

func (m *exampleRepository) Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error) {
	data := new(datastruct.ExampleDetail)

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail from example e
	where e.uuid = '%v'
	`, *id)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return data, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, &util.Error{
			Errors:     err.Error(),
			StatusCode: 400,
			Message:    infrastructure.Localize("NOT_FOUND"),
		}
	}

	return data, &util.Error{}
}

func (m *exampleRepository) Delete(ctx *context.Context, id *string) *util.Error {
	sqlRslt, err := m.sqlxDB.ExecContext(*ctx, fmt.Sprintf("delete from example where uuid = '%v'", *id))
	if err != nil {
		return &util.Error{
			Errors: err.Error(),
		}
	}

	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     "no data",
			Message:    infrastructure.Localize("FAILED_DELETE_NO_DATA"),
			StatusCode: 400,
		}
	}

	return &util.Error{}
}
