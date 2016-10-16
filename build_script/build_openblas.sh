#!/bin/bash

script_path=$(cd "$(dirname "$0")"; pwd)
third_party_path="${script_path}/../third_party"
target_path="${third_party_path}/target"
openblas_path="${third_party_path}/OpenBLAS"

echo -e "${third_party_path}"

mkdir -p $target_path

os_name=`uname`

pushd $openblas_path
make OSNAME=$os_name BINARY=64
make install PREFIX=$target_path
popd
