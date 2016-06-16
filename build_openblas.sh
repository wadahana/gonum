#!/bin/bash

OPENBLAS_DIR=$PWD/blas
SOURCE_DIR=$OPENBLAS_DIR/OpenBLAS
TARGET_DIR=$OPENBLAS_DIR/target
OS_NAME=`uname`
pushd blas

mkdir -p $TARGET_DIR

pushd OpenBLAS

make OSNAME=$OS_NAME BINARY=64
make install PREFIX=$TARGET_DIR
