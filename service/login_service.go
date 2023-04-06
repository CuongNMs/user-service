package service

import (
	"context"
	"user-service/common"
	"user-service/model"
)

type UserRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error)
}

type loginService struct {
	repo UserRepo
}

func NewLoginService(repo UserRepo) *loginService {
	return &loginService{
		repo: repo,
	}
}

func (service *loginService) Login(ctx context.Context, data *model.UserLogin) (*common.Token, error) {
	user, err := service.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.NewErrorResponse(401, "Login error", err)
	}
	hash := common.NewMd5Hash()
	passHashed := hash.Hash(data.Password + user.Salt)
	if passHashed != user.Password {
		return nil, common.LoginErrorResponse(403, "Login error")
	}
	return common.NewToken().GenerateToken(user.FirstName, user.LastName), nil
}
