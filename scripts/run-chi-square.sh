#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

# rm -rf out/chi-square/ || echo "Skipped"
rm out-chi-square.json || echo "Skipped"

mkdir -p out/chi-square/

cd src/chi-square

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
source .env/bin/activate || . .env/bin/activate

# pip install pandas statsmodels
# pip freeze > requirements.txt
pip install -r requirements.txt

python main.py > ../../out/out-chi-square.json 2> ../../out/out-chi-square.json
