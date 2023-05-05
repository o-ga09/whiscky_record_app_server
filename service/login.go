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
	_, err := l.Repo.GetUser(ctx,l.DB,uid)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w",err)
	}

	return "ok",nil
}