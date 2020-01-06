#!/bin/bash

set -x

while getopts "a:v:h" opt; do
  case $opt in
	a) 
	ARCH=$OPTARG
	;;
	v)
	VERSION=$OPTARG
	;;
	h)
      cat <<EOF
Usage: $0 [-a build_arch] [-v version]
Support Arch: windows, linux
EOF
      exit 0
      ;;
  esac
done

if [ "$ARCH"x = "x" ];then
	ARCH=linux
fi

if [ "$VERSION"x = "x" ];then
	VERSION=latest
fi


# build
go mod vendor
GOOS=$ARCH GOARCH=amd64 go build --ldflags '-w -s'

if [ -f "appLegu" ];then
	cp appLegu bin/
elif [ -f "appLegu.exe" ];then
	cp appLegu.exe bin/
else
	echo "appLegu build artificat not found , run `go build` for debug."
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
	upx bin/appLegu*
fi

chmod u+x applegu/lib/zipalign
chmod u+x applegu/bin/appLegu

if [ "$ARCH"x = "windowsx" ];then
	tar vczf applegu-windows-${VERSION}.tar.gz applegu
else
	tar vczf applegu-${VERSION}.tar.gz applegu
fi

# clean
rm -f appLegu
rm -f appLegu.exe
rm -f bin/*
rm -rf applegu


# install
# tar vxzf applegu-latest.tar.gz -C /usr/local/

# run 
# /usr/local/applegu/bin/appLegu
