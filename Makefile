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
	cd bin && rm *
	cd cmd/grubenlampe-server && go build -o ../../bin/grubenlamped
	cd cmd/grubenlampe && go build -o ../../bin/grubenlampe

server:
	cd cmd/grubenlampe-server && go build -o ../../bin/grubenlamped

client:
	cd cmd/grubenlampe && go build -o ../../bin/grubenlampe

