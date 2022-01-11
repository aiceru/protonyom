#!/bin/bash

SRC_DIR=$(pwd)
DART_OUT=$SRC_DIR/dart
GO_OUT=$SRC_DIR
GO_DIR=$SRC_DIR/goprotonyom

rm -rf $GO_DIR
rm -rf $DART_OUT
mkdir -p $GO_DIR
mkdir -p $DART_OUT

protoc \
-I=$SRC_DIR \
--dart_out=$DART_OUT \
--go_out=$GO_OUT \
--go_opt=module=github.com/aiceru/protonyom \
$SRC_DIR/ohmnyom.proto

if [ "$(ls -A $GO_DIR)" ]; then
  cd $GO_DIR
  go mod init github.com/aiceru/protonyom/goprotonyom
  go mod tidy
  cd ..
fi