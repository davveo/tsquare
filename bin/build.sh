#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin go build  -o user-srv /Users/davve/Work/Go/tsquare/srv/user-srv/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o conf-srv /Users/davve/Work/Go/tsquare/srv/conf-srv/*.go
CGO_ENABLED=0 GOOS=darwin go build  -o auth-srv /Users/davve/Work/Go/tsquare/srv/auth-srv/*.go
