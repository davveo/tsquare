package basic

import (
	"github.com/zbrechave/tsquare/basic/config"
	"github.com/zbrechave/tsquare/basic/db"
	"github.com/zbrechave/tsquare/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
