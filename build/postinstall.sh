#! /bin/bash

DIR=$(pwd)

# Patch wiredep so we can use it to manage NPM dependencies instead of bower
cd ./node_modules/wiredep
patch -N < ../../build/patch/wiredep/wiredep.patch
cd lib
patch -N < ../../../build/patch/wiredep/detect-dependencies.patch
cd ..
rm lib/*.orig lib/*.rej *.orig *.rej 2> /dev/null

cd ${DIR}

# Govendor is required by the project. Install it in the .tools directory.
GOPATH=`pwd`/.tools/go go get github.com/kardianos/govendor
# XtbGeneator is required by the project. Clone it into .tools.
if ! [ -a "./.tools/xtbgenerator/bin/XtbGenerator.jar" ]
then
  (cd ./.tools/; git clone https://github.com/kuzmisin/xtbgenerator; cd xtbgenerator; git checkout d6a6c9ed0833f461508351a80bc36854bc5509b2)
fi