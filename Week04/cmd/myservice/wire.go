package main

import (
	"github.com/google/wire"
	"week04/internal/myservice/biz"
	"week04/internal/myservice/data"
)

//go:generate wire
func NewUserUsecase() (*biz.UserUsecase, func(), error) {
	//userSet:=wire.NewSet(data.NewUserRepo,biz.NewUserUsecase)
	panic(wire.Build(biz.NewUserUsecase, data.NewUserRepo, data.DbSet))
	//return biz.UserUsecase{}
}
