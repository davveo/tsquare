#!/usr/bin/env bash

./conf-srv
./auth-srv --registry=etcd
./user-srv --registry=etcd
