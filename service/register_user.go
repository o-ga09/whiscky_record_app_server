package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"main/auth"
	"main/entity"
	"main/store"

	"github.com/google/uuid"
)

type RegisterUser struct {
	DB store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, user_id string) (*entity.User, error) {
	token,err := auth.GetUserInfo(ctx,user_id)
	if err != nil {
		return nil,fmt.Errorf("cannot get user info: %w",err)
	}

	uid := base64.StdEncoding.EncodeToString([]byte(token))
	uuid, err := uuid.Parse(uid)
	if err != err {
		return nil,err
	}
	u := &entity.User{
		User_ID: entity.UserID(uuid),
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register: %w",err)
	}
	return u, nil
}