package handler

import (
	"encoding/json"
	"main/entity"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RecordWhicky struct {
	Service RecordWhickyService
	Validator *validator.Validate
}

func (rw *RecordWhicky) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	ctx := r.Context()
	var b struct {
		Uid entity.UserID `json:"uid" validate:"require"`
		Name string `json:"name" validate:"require"`
		ImageURL string `json:"image" validate:"require"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	if err := rw.Validator.Struct(b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusBadRequest)
	}

	status, err := rw.Service.RecordWhicky(ctx,b.Uid,b.Name,b.ImageURL)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}

	rsp := struct {
		Status string `json:"status"`
	}{Status: status}
	RespondJSON(ctx,w,rsp,http.StatusOK)
}