#!/usr/bin/env bash

micro --registry=etcd  api --handler=http

../build/userApi  --registry=etcd
../build/confSrv  --registry=etcd
../build/authSrv  --registry=etcd
../build/userSrv  --registry=etcd
