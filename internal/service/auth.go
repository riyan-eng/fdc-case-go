package service

import (
	"context"
	"database/sql"
	"server/config"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/util"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx *context.Context, ent *entity.AuthRegister) *util.Error
	Login(ctx *context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error)
	Refresh(ctx *context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, *util.Error)
	Me(ctx *context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, *util.Error)
	Logout(ctx *context.Context, ent *entity.AuthLogout) *util.Error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao *repository.DAO) AuthService {
	return &authService{
		dao: *dao,
	}
}

func (m *authService) Register(ctx *context.Context, ent *entity.AuthRegister) *util.Error {
	hashPassword, errT := util.GenerateHash(ent.Password)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}

	modelUser := model.User{
		Id:       *ent.UserId,
		Username: sql.NullString{String: *ent.UserName, Valid: util.NewIsValid().String(ent.UserName)},
		Password: sql.NullString{String: hashPassword, Valid: true},
		IsActive: sql.NullBool{Bool: true, Valid: true},
	}

	modelUserData := model.UserData{
		Id:       uuid.NewString(),
		UserId:   sql.NullString{String: modelUser.Id, Valid: true},
		RoleCode: sql.NullString{String: *ent.RoleCode, Valid: true},
	}
	if err := m.dao.NewAuthRepository().Register(ctx, &modelUser, &modelUserData); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *authService) Login(ctx *context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error) {
	token := new(datastruct.AuthToken)

	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Username)
	if err.Errors != nil {
		return data, token, err
	}

	if !data.IsActive {
		return data, token, &util.Error{
			Errors:     "user",
			Message:    "user tidak aktif",
			StatusCode: 400,
		}
	}

	// verify password
	if !util.VerifyHash(data.Password, *ent.Password) {
		return data, token, &util.Error{
			Errors:     "password",
			Message:    "Username atau Password yang anda inputkan salah, silakan menginputkan username dan password yang benar",
			StatusCode: 400,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors: errT.Error(),
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors: errT.Error(),
		}
	}
	enforce := config.NewEnforcer()
	enforce.AddRoleForUser(data.Id, data.RoleCode)
	return data, &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Refresh(ctx *context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, *util.Error) {
	newRefresh := new(datastruct.AuthToken)
	claim, errT := util.NewToken().ParseRefresh(ent.Token)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT.Error(),
			StatusCode: 401,
		}
	}

	if errT := util.NewToken().ValidateRefresh(ctx, claim); errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT.Error(),
			StatusCode: 401,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors: errT.Error(),
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors: errT.Error(),
		}
	}

	return &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Logout(ctx *context.Context, ent *entity.AuthLogout) *util.Error {
	err := m.dao.NewAuthRepository().Logout(ctx, ent.UserId)
	if err.Errors != nil {
		return err
	}

	return &util.Error{}
}

func (m *authService) Me(ctx *context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, *util.Error) {
	data, err := m.dao.NewAuthRepository().Me(ctx, ent.UserId)
	if err.Errors != nil {
		return data, err
	}

	return data, &util.Error{}
}
