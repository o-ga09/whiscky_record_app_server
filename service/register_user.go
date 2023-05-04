package service

import (
	"context"
	"fmt"

	"main/entity"
	"main/store"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, user_id string) (*entity.User, error) {
	uid ,err := bcrypt.GenerateFromPassword([]byte(user_id), bcrypt.DefaultCost)
	if err != nil {
		return nil,fmt.Errorf("cannot hash password: %w",err)
	}
	u := &entity.User{
		User_ID: string(uid),
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register: %w",err)
	}
	return u, nil
}