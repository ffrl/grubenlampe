SHELL = /bin/bash
GO = go
DIR = $(shell pwd)

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m


clean:
	cd bin && rm *

all:
	cd cmd/grubenlampe-server && go build -o ../../bin/grubenlamped
	cd cmd/gl && go build -o ../../bin/gl

server:
	cd cmd/grubenlampe-server && go build -o ../../bin/grubenlamped

client:
	cd cmd/gl && go build -o ../../bin/gl

