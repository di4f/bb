#!/bin/sh

wd=`pwd`

cd $wd/src/cmd/goblin && go install && cd $wd

mkdir -p $HOME/app/goblin
cp -rf $wd/app/* $HOME/app/goblin

