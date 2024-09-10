#!/usr/bin/env bash

REPOSITORY=$1

WITH_BUILD=${2:-false}

SCRIPT_BASE_PATH=`pwd`
echo "SCRIPT_BASE_PATH: $SCRIPT_BASE_PATH"

source ./extract-infos.sh $REPOSITORY $WITH_BUILD $SCRIPT_BASE_PATH || . ./extract-infos.sh $REPOSITORY $WITH_BUILD $SCRIPT_BASE_PATH

