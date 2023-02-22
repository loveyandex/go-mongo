package service

import (
	"github.com/loveyandex/go-mongo/db"
	"github.com/loveyandex/go-mongo/model"
)


type P2PSrv struct {
	CmnSrv    *CmnSrv[model.User]
	
}

func NewP2PSrv() *P2PSrv {
	return &P2PSrv{CmnSrv: &CmnSrv[model.User]{Xcol: db.Collection("p2ps")}}
}