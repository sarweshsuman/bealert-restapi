package resources

import (
	"github.com/sarweshsuman/bealert-restapi/redis"
	"github.com/robfig/config"
	"github.com/sarwesh/utilities"	
	"strconv"
	"errors"
)

type Cell struct {
	Id string `json:"id"`
	Msisdns []int `json:"msisdns"`
}

func GetCell(id string,cfg *config.Config) (*Cell,error) {
	connection , err  := redis.New(cfg)
	if err != nil {
		return &Cell{},err
	}
	result , err := connection.HGetAll(utilities.ConcatStrings("cell:",id))
	if err != nil {
		return &Cell{},err
	}
	if len(result) == 0 {
		return &Cell{},errors.New("No msisdn Found")
	}
	cell := new(Cell)
	cell.Id=id
	for key,_ := range result {
		msisdn , err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		cell.Msisdns = append(cell.Msisdns,msisdn)
	}
	return cell,nil
}