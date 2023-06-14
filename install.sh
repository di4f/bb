#!/bin/sh

wd=`pwd`

# cd $wd/src/cmd/goblin && go install && cd $wd
go install

mkdir -p $HOME/app/goblin
cp -rf $wd/app/* $HOME/app/goblin

