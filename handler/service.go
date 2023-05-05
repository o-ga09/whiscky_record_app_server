package handler

import (
	"context"

	"main/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . RegisterUserService LoginService
type RegisterUserService interface {
	RegisterUser(ctx context.Context, uid string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, uid string) (string, error)
}

type RecordWhickyService interface {
	RecordWhicky(ctx context.Context,uid string, name string,imageurl string) (string,error)
}

type GetWhickyService interface {
	GetRecord(ctx context.Context,uid string) (*[]entity.WhickyRecord,error)
}