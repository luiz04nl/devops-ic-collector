#!/usr/bin/env bash

echo "The first argument is: $1"
REPOSITORY=$1
echo "REPOSITORY: $1"

cd ../../repos
cd $REPOSITORY
# git submodule update --init --recursive

MSYS_NO_PATHCONV=1 docker run --rm -v $(pwd)/../../repos/:/repos/ ck $REPOSITORY 2>&1

PROJECT_OUT="../../repos/$REPOSITORY/ck-out"

mkdir -p ../../out/code-metrics/$REPOSITORY

cp -rvf $PROJECT_OUT/* ../../out/code-metrics/$REPOSITORY/

echo "Finished"