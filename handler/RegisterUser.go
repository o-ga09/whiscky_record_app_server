package handler

import (
	"encoding/json"
	"main/entity"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Service RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Userid string `json:"user_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	if err := ru.Validator.Struct(b); err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusBadRequest)
	}

	u, err := ru.Service.RegisterUser(ctx,b.Userid)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.UserID `json:"user_id"`
	}{ID: u.ID}
	RespondJSON(ctx,w,rsp,http.StatusOK)
}