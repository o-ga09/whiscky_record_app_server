package service

import (
	"context"
	"database/sql"
	"errors"
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
	if errors.Is(err,sql.ErrNoRows) {
		return "",err
	} else if err != nil {
		return "", fmt.Errorf("failed to list: %w",err)
	}

	return fmt.Sprint(u.User_ID),nil
}