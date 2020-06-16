package redis

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var Config *ini.File
var CsrfExcept *ini.File
var RootPath string

func init()  {
	RootPath="/Users/davve/Work/Go/tsquare/config"
	var err error
	Config, err = ini.Load(RootPath+"/config.ini");
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
