package main

import (
	"context"
	"net/http"

	"main/clock"
	"main/config"

	"main/service"
	"main/store"

	"main/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", HealthCheckHandler)
	v := validator.New()
	db, cleanup, err := store.New(ctx,cfg)
	if err != nil {
		return nil, cleanup, err
	}
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}

	l := &handler.Login{
		Service: &service.Login{
			DB: db,
			Repo: &r,
		},
		Validator: v,
	}

	mux.Post("/login",l.ServeHTTP)

	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register",ru.ServeHTTP)
	return mux, cleanup, nil
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}