#!/bin/bash

path=$(cd "$(dirname "$0")"; pwd)
third_party_path=$path

target_path="${third_party_path}/target"
openblas_path="${third_party_path}/OpenBLAS"
alglib_path="${third_party_path}/alglib.3.10.0"

echo -e "${third_party_path}"

mkdir -p "${target_path}"



echo -e "\n\t\tbuild OpenBLAS\n\n"

pushd "${openblas_path}" > /dev/null
os_name=`uname`
make OSNAME=$os_name BINARY=64
make install PREFIX="${target_path}"
popd > /dev/null

echo -e "\n\t\tbuild alglib\n\n"

pushd "${alglib_path}" > /dev/null
make
make install
popd  > /dev/null