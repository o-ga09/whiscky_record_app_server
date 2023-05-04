package service

import (
	"context"
	"fmt"

	"main/store"
)

type Login struct {
	DB store.Queryer
	Repo UserGetter
	TokenGenerator TokenGenerator
}

func (l *Login) Login(ctx context.Context, uid string) (string,error) {
	u, err := l.Repo.GetUser(ctx,l.DB,uid)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w",err)
	}
	if err := u.ComparePassword(uid); err != nil {
		return "", fmt.Errorf("wrong user id: %w",err)
	}

	return "ok",nil
}