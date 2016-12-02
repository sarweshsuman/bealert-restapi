package redis

import (
	"time"
	"errors"
)

func (connection *RedisConnection) Keys(pattern string) ([]string , error ) {
	cmd := connection.Client.Keys(pattern)
	result , err := cmd.Result()
	if connection.checkError(err) == true {
		return []string{},err
	}
	return result,nil
}


func (connection *RedisConnection) Set(key string,value interface{}) error {
	cmd := connection.Client.Set(key,value,time.Duration(0))
	_ , err := cmd.Result()
	if connection.checkError(err) == true {
		return err
	}
	return nil
}

func (connection *RedisConnection) HSet(key ,field , value string) error {
	cmd := connection.Client.HSet(key,field,value)
	isItSet , err := cmd.Result()
	if connection.checkError(err) == true {
		return err
	}
	if isItSet == false {
		connection.Log("Unable to set the field, May be no new value provided? ")
		return errors.New("Unable to set the field, May be no new value provided? ")
	}
	return nil	
}

func (connection *RedisConnection) HGetAll(key string) (map[string]string , error ) {
	cmd := connection.Client.HGetAll(key)
	result, err := cmd.Result()
	if connection.checkError(err) == true {
		return map[string]string{},err
	}
	return result,nil
}