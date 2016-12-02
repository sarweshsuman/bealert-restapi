package handlers

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/sarweshsuman/bealert-restapi/resources"
	"encoding/json"
)

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w,`<html><body>
		<h1>Welcome</h1>
		<hr/>
		<p>This is a Rest API to fetch cell or imsi info</p>
		<p>Two APIs are exposed</p>
		<ul>
			<li>/bealert/cell/{cellid:lac}</li>
			<li>/bealert/imsi/{imsi}</li>
		</ul>
		</body>
		</html>`)
}

func (Cfg *Config)  fetchCellInfo(w http.ResponseWriter, r *http.Request, urlparam httprouter.Params){
	id := urlparam.ByName("id")
	cell, err := resources.GetCell(id,Cfg.cfg)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type","application/html")				
		fmt.Fprintf(w,"Internal Server Error ",err.Error())
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(cell)
}

func (Cfg *Config)  fetchImsiInfo(w http.ResponseWriter, r *http.Request, urlparam httprouter.Params){
	id := urlparam.ByName("id")
	imsi, err := resources.GetImsi(id,Cfg.cfg)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)		
		w.Header().Set("Content-Type","application/html")		
		fmt.Fprintf(w,"Internal Server Error ",err.Error())
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(imsi)
}