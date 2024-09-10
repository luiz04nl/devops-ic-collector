#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

# rm -rf out/regression/ || echo "Skipped"
rm out-regression.json || echo "Skipped"

mkdir -p out/regression/

cd src/regression

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
source .env/bin/activate || . .env/bin/activate

# pip install pandas statsmodels
# pip freeze > requirements.txt
pip install -r requirements.txt

python main.py > ../../out/out-regression.json 2> ../../out/out-regression.json
