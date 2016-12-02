package handlers

import (
	"github.com/julienschmidt/httprouter"
)

func (Cfg *Config) GetRouter() *httprouter.Router{
	router := httprouter.New()
	router.GET("/",about)
	router.GET("/bealert/cell/:id",Cfg.fetchCellInfo)
	router.GET("/bealert/imsi/:id",Cfg.fetchImsiInfo)
	return router
}