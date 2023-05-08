package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"main/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Login struct {
	Service LoginService 
	Validator *validator.Validate
}

func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		User_ID string `json:"token" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx,w,&ErrResponse {
			Message: err.Error(),
		},http.StatusInternalServerError)
		log.Printf("failed json decode")
		return
	}
	err := l.Validator.Struct(body)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusBadRequest)
		log.Printf("failed json validate")
		return
	}

	decode_str,err := base64.StdEncoding.DecodeString(body.User_ID)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse {
			Message: err.Error(),
		},http.StatusInternalServerError)
		log.Printf("failed decode")
		return
	}

	token,err := auth.GetUserInfo(ctx,string(decode_str))
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse {
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}

	uid := base64.StdEncoding.EncodeToString([]byte(token))
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse {
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}

	id, err := l.Service.Login(ctx,uid)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		log.Printf("failed login authrazation")
		return
	}
	rsp := struct {
		Id string `json:"user_id" validate:"required"`
	}{
		Id: id,
	}

	RespondJSON(ctx,w,rsp,http.StatusOK)

}