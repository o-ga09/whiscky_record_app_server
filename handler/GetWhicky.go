package handler

import (
	"encoding/json"
	"main/entity"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type GetRecord struct {
	Service GetWhickyService
	Validator *validator.Validate
}

func (gw *GetRecord) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Uid entity.UserID `json:"uid" validate:"require"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	if err := gw.Validator.Struct(b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusBadRequest)
	}

	record,err := gw.Service.GetRecord(ctx,b.Uid)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	rsp := struct {
		Result *entity.WhickyRecord `json:"result"`
	}{Result: record}
	RespondJSON(ctx,w,rsp,http.StatusOK)
}