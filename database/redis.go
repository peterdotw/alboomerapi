package database

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func initRedis() redis.Conn {
	connection, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Redis up and running")
	return connection
}

// RedisConnection -  Redis initialized
var RedisConnection = initRedis()
