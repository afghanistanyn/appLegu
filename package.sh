#!/bin/bash

set -e

# build 
go mod vendor
go build --ldflags '-w -s'


if [ -f "appLegu" ];then
	cp appLegu bin/
else
	echo "appLegu not found , run 'go build before package'"
	exit -1
fi

mkdir -p applegu

cp -rf bin/ applegu
rm -f applegu/bin/.gitkeep
cp -rf conf/ applegu

cp -rf lib/ applegu
cp -rf pkgs/ applegu
rm -f applegu/pkgs/.gitkeep


if [ "$(hash upx)x" = "x" ];then
	upx bin/appLegu
fi

tar vczf applegu.tar.gz applegu

# clean
rm -f appLegu
rm -rf applegu


# install
# tar vxzf applegu.tar.gz -C /usr/local/
# chmod u+x /usr/local/applegu/lib/zipalign

# run 
# /usr/local/applegu/bin/appLegu
