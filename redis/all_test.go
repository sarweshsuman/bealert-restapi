package redis

import (
	"testing"
	"github.com/robfig/config"
	"fmt"
)

func createDummyConfiguration() (*config.Config) {

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

func TestNew(t *testing.T){
	// Creating a dummy configuration
	c := config.NewDefault()
	_ , error := New(c)

	if error == nil {
		t.Error("Failed to verify configuration")
	}	
	c.AddSection("redis")
	c.AddOption("redis","port","6379")
	c.AddOption("redis","timeout","50000")
	c.AddOption("redis","db","0")
	c.AddOption("redis","password","")

	conn , error := New(c)

	if error == nil {
		t.Fatal("Failed to identify missing parameters")
	} 
	
	c.AddOption("redis","host","localhost")
	conn , error = New(c)

	if error != nil {
		t.Fatal(error)
	} 

	_ , error  = conn.Client.Ping().Result()

	if error != nil {
		t.Fatal("Failed to connect to localhost:6379 database:0" , error)
	} 
}

func TestSet(t *testing.T){

	c := createDummyConfiguration()

	conn , err := New(c)

	if err != nil {
		t.Fatal("Failed to connect to redis ",err)
	}

	err = conn.Set("sarwesh","suman")

	if err != nil {
		t.Error("Failed to set a value ",err)
	}

}


func TestKeys(t *testing.T){
	
	c := createDummyConfiguration()

	conn , err := New(c)

	if err != nil {
		t.Fatal("Failed to connect to redis ",err)
	}

	err = conn.Set("sarwesh","suman")

	if err != nil {
		t.Error("Failed to set a value ",err)
	}

	listOfKeys , err := conn.Keys("sarwes*")

	if len(listOfKeys) == 0  || err != nil {
		t.Error("Failed to reterieve list of keys based on pattern ",err)
	}


}

func TestHSet(t *testing.T){
	
	c := createDummyConfiguration()

	conn , err := New(c)

	if err != nil {
		t.Fatal("Failed to connect to redis ",err)
	}

	err = conn.HSet("sarwesh","suman","suman")

	if err == nil {
		t.Error("Failed to capture error it should have thrown an eror since key sarwesh is a simple key-value")
	}

	err = conn.HSet("sarweshhash","suman","suman1")

	if err != nil {
		t.Error("Failed in setting value for a field within a Hash ",err)
	}

}

func TestHGetAll(t *testing.T){
	
	c := createDummyConfiguration()

	conn , err := New(c)

	if err != nil {
		t.Fatal("Failed to connect to redis ",err)
	}

	conn.HSet("sarweshhash","suman1","suman1")
	conn.HSet("sarweshhash","suman2","suman2")
	conn.HSet("sarweshhash","suman3","suman3")
	conn.HSet("sarweshhash","suman4","suman4")
	conn.HSet("sarweshhash","suman5","suman5")

	maps , err := conn.HGetAll("sarweshhash")

	if err != nil {
		t.Error("Failed in getting all the values for a hash ",err)
	} else {
		fmt.Println(maps)
	}


}
