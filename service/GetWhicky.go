package service

import (
	"context"
	"fmt"
	"main/entity"
	"main/store"
)

type GetRecord struct {
	DB store.Queryer
	Repo RecordGetter
}

func (gw *GetRecord) GetRecord(ctx context.Context,uid entity.UserID) (*entity.WhickyRecord,error) {
	record,err := gw.Repo.GetWhickyRecord(ctx,gw.DB,uid)
	if err != nil {
		return nil,fmt.Errorf("faild to get record")
	}
	return record,nil
}