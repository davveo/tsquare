package basic

import (
	"auth-srv/basic/config"
	"auth-srv/basic/db"
	"auth-srv/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
