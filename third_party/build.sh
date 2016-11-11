#!/bin/bash

path=$(cd "$(dirname "$0")"; pwd)
third_party_path=$path

target_path="${third_party_path}/target"
openblas_path="${third_party_path}/OpenBLAS"
alglib_path="${third_party_path}/alglib.3.10.0"

echo -e "${third_party_path}"

mkdir -p "${target_path}"


function build_openblas()
{
    echo -e "- - - - - - - - - - - - - - - - - - - - - - \n"
    echo -e "\n\t\tbuild OpenBLAS\n\n"
    echo -e "- - - - - - - - - - - - - - - - - - - - - - \n"
    openblas_github_url="https://github.com/xianyi/OpenBLAS.git"

    echo -e "chekcout OpenBLAS.git ...\n"
    if [ ! -d "$openblas_path/.git" ]; then
        git clone "$openblas_github_url" "$openblas_path"  > /dev/null
    fi
    pushd "$openblas_path" > /dev/null
    git checkout
    git reset --hard HEAD
    if [ ! -f "Makefile" ]; then
        echo -e "git clone ${openblas_github_url} fail..\n"
        exit 0
    fi

    os_name=`uname`
    make OSNAME=$os_name BINARY=64
    make install PREFIX="${target_path}"
    popd > /dev/null
}

function build_alglib()
{
    echo -e "- - - - - - - - - - - - - - - - - - - - - - \n"
    echo -e "\n\t\tbuild alglib\n\n"
    echo -e "- - - - - - - - - - - - - - - - - - - - - - \n"

    pushd "${alglib_path}" > /dev/null
    make
    make install
    popd  > /dev/null
}

function usage()
{
    echo -e "build.sh  command"
    echo -e "command: "
    echo -e "          openblas                build OpenBLAS"
    echo -e "          alglib                  build alglib.3.10.0"
    echo -e "          all                     build OpenBLAS, alglib"
}

if [ $# != 1 ] ; then
    usage
    echo -e "\a\a\a"
    exit 0
fi

if [ $1 == "all" ] ; then
    build_openblas
    build_alglib
elif [ $1 == "openblas" ] ; then
    build_openblas
elif [ $1 == 'alglib' ] ; then
    build_alglib
else
    usage
fi

