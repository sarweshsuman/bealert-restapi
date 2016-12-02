package redis

import (
	"log"
	redis "gopkg.in/redis.v5"
	"github.com/robfig/config"
	"github.com/sarwesh/utilities"
	"errors"
	"strconv"
)

// This go code is responsible for connecting with redis instance.
// It expects configuration from config package

type RedisConnection struct {
	Host string
	Port int
	Timeout int
	Database int
	Password string
	Client *redis.Client
	Log func(string)
}

func New(cfg *config.Config) (*RedisConnection,error) {
	doILog := false
	if cfg.HasOption("log","enabled") == true {
		doILog , _ =  cfg.Bool("log","enabled")
	}
	isAllParametersFound := cfg.HasOption("redis","host") == true && cfg.HasOption("redis","port") == true && cfg.HasOption("redis","timeout") == true && cfg.HasOption("redis","db") == true && cfg.HasOption("redis","password") == true
	if isAllParametersFound == true {
				if doILog == true {
					log.Println("All Parameters found for Redis Connection")
				}
				redisConnection := new(RedisConnection)
				redisConnection.Host, _ = cfg.String("redis","host")
				redisConnection.Port, _ = cfg.Int("redis","port")
				redisConnection.Timeout, _ = cfg.Int("redis","timeout")
				redisConnection.Database, _ = cfg.Int("redis","db")
				redisConnection.Password, _ = cfg.String("redis","password")
				redisConnection.Log = func(line string){
					if doILog == true {
						log.Printf(line)
					}
				}
				redisConnection.Client = redis.NewClient(&redis.Options{
					Addr: utilities.ConcatStrings(redisConnection.Host,":",strconv.Itoa(redisConnection.Port)),
					Password: redisConnection.Password,
					DB: redisConnection.Database,
					})

				return redisConnection,nil
		} else {
				if doILog == true {
					log.Println("All mandatory parameters not found in configuration")
				}
			return &RedisConnection{},errors.New("All mandatory parameters not found in configuration")
		}
}

func (connection *RedisConnection) checkError(err error) bool{
	if err != nil {
		connection.Log(err.Error())		
		return true
	}
	return false
}