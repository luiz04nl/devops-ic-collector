#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

mkdir -p out/sample-distribution

cd src/sample-distribution

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
. .env/bin/activate
pip install -r requirements.txt

python main.py > ../../out/out-sample-distribution.json 2> ../../out/out-sample-distribution.json