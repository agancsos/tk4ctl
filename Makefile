###############################################################################
# Name        : Makefile                                                      #
# Author      : Abel Gancsos                                                  #
# Version     : v. 1.0.0.0                                                    #
# Description : Helps build the utility.                                      #
###############################################################################
.DEFAULT_GOAL := build

build:
	which go
	if [ ! -d "./dist" ]; then mkdir ./dist; fi 
	GOPATH=$(PWD)/packages go build -o ./dist/tk4ctl ./src/main.go

clean:
	if [ -d "./dist" ]; then rm -fr ./dist/**; fi		
