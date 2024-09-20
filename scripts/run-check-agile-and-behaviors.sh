#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

pythonOutputDirectory="out/check-agile-and-behaviors"
if [ -d "$pythonOutputDirectory" ]; then
    rm -r $pythonOutputDirectory
fi
mkdir -p $pythonOutputDirectory

cd src/check-agile-and-behaviors

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
. .env/bin/activate
pip install -r requirements.txt

go run . > ../../out/out-check-agile-and-behaviors.json 2> ../../out/out-check-agile-and-behaviors.json
