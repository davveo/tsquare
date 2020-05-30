#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin go build  -o ../build/userApi /Users/davve/Work/Go/tsquare/api/user/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o ../build/userSrv /Users/davve/Work/Go/tsquare/srv/user-srv/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o ../build/confSrv /Users/davve/Work/Go/tsquare/srv/conf-srv/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o ../build/authSrv /Users/davve/Work/Go/tsquare/srv/auth-srv/*.go
