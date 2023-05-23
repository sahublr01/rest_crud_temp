USER=$(shell id -u -n)
COMMIT=$(shell git log --format="%H" -n 1)
TIME=$(shell date)
BASE=gitlab.dell.com/Aman_Sharma/upgradeM3u8/

#version 
VERSION=v1.0.0
#binary name
OUTPUT=upgradeM3u8

build:
	@go build  -o ${OUTPUT} -v -ldflags="-X '${BASE}config.Build_Version=${VERSION}' -X '${BASE}config.Build_User=${USER}' -X '${BASE}config.Build_Time=${TIME}' -X '${BASE}config.Build_Commit=${COMMIT}'"
	@echo Build Generated