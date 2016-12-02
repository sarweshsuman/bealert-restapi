package resources 

import (
	"testing"
	"github.com/robfig/config"	
	"fmt"
)

func createDummyConfiguration() *config.Config{

	c := config.NewDefault()

	c.AddSection("redis")
	c.AddOption("redis","host","localhost")	
	c.AddOption("redis","port","6379")
	c.AddOption("redis","timeout","50000")
	c.AddOption("redis","db","0")
	c.AddOption("redis","password","")
	c.AddOption("log","enabled","true")

	return c
}

func TestGetCell(t *testing.T){
	c  := createDummyConfiguration()
	
	result , err := GetCell("190:200",c)
	if  err != nil {
		t.Error("Failed to Get cell info ", err)
	}
	if result.Id == "" {
		t.Error("Empty result")
	}
}

func TestGetImsi(t *testing.T){
	c  := createDummyConfiguration()

	result , err := GetImsi("987654321",c)
	
	if  err != nil {
		t.Error("Failed to Get cell info ", err)
	}
	if result.Imsi == "" {
		t.Error("Empty result")
	}	
}