package service

import (
	"context"
	"fmt"
	"main/entity"
	"main/store"
)

type RecordWhicky struct {
	DB store.Execer
	Repo WhickyRegister
}

func (rw *RecordWhicky) RecordWhicky(ctx context.Context,uid string,name string,evaluate string, imageURL string) (string,error) {
	record := &entity.WhickyRecord{
		UserID: uid,
		Name: name,
		Evaluate: evaluate,
		ImageURL: imageURL,
	}
	if err := rw.Repo.RegisterWhicky(ctx,rw.DB,record);err != nil {
		return "",fmt.Errorf("faild to register: %w",err)
	}
	return "ok",nil
}