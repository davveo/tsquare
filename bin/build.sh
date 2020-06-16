#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin go build  -o ../build/userApi /Users/davve/Work/Go/tsquare/api/user/*.go
#CGO_ENABLED=0 GOOS=darwin go build  -o ../build/userSrv /Users/davve/Work/Go/tsquare/srv/user/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o ../build/confSrv /Users/davve/Work/Go/tsquare/srv/conf/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o ../build/smsSrv /Users/davve/Work/Go/tsquare/srv/sms/*.go
#CGO_ENABLED=0 GOOS=darwin go build  -o ../build/confSrv /Users/davve/Work/Go/tsquare/srv/uuid/*.go
#CGO_ENABLED=0 GOOS=darwin go build  -o ../build/authSrv /Users/davve/Work/Go/tsquare/srv/auth/*.go
