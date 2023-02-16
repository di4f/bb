#!/bin/sh

wd=`pwd`

cd src/cmd/goblin && go install && cd $wd

mkdir -p $HOME/app/goblin
cp -rf app/* $HOME/app/goblin

