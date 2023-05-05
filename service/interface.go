package service

import (
	"context"

	"main/entity"
	"main/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . UserRegister UserGetter TokenGenerator
type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}

type WhickyRegister interface {
	RegisterWhicky(ctx context.Context,db store.Execer,record *entity.WhickyRecord) error
}

type RecordGetter interface {
	GetWhickyRecord(ctx context.Context,db store.Queryer,uid entity.UserID) (*entity.WhickyRecord,error)
}