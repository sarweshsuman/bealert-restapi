package resources 

import (
	"github.com/sarweshsuman/bealert-restapi/redis"
	"github.com/robfig/config"
	"github.com/sarwesh/utilities"
	"errors"
)

type Imsi struct {
	Imsi string `json:"imsi"`
	Msisdn string `json:"msisdn"`
	Cellid string `json:"cellid"`
	NetworkEventTs string `json:"network_event_ts"`
}

func GetImsi(imsi string,cfg *config.Config) (*Imsi,error) {
	connection , err  := redis.New(cfg)
	if err != nil {
		return &Imsi{},err
	}
	result , err := connection.HGetAll(utilities.ConcatStrings("imsi:",imsi))
	if err != nil {
		connection.Log(err.Error())
		return &Imsi{},err
	}
	if len(result) == 0 {
		return &Imsi{},errors.New("No detail for IMSI found")
	}
	imsii := new(Imsi)	
	imsii.Imsi=imsi
	imsii.Msisdn=result["msisdn"]
	imsii.Cellid=result["cellid"]	
	imsii.NetworkEventTs=result["network_event_ts"]
	connection.Log("Data fetch successfull")
	return imsii,nil
}