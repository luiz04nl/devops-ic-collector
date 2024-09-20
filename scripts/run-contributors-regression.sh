#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

# rm -rf out/contributors-regression/ || echo "Skipped"
rm out-contributors-regression.json || echo "Skipped"

mkdir -p out/contributors-regression/

cd src/contributors-regression

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
. .env/bin/activate

# pip install pandas statsmodels
# pip freeze > requirements.txt
pip install -r requirements.txt

python main.py > ../../out/out-contributors-regression.json 2> ../../out/out-contributors-regression.json
