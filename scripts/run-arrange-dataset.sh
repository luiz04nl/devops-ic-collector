#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

rm -rf out/arrange-dataset/ || echo "Skipped"
rm out-arrange-dataset.json || echo "Skipped"

mkdir -p out/arrange-dataset/

cd src/arrange-dataset

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
. .env/bin/activate

# pip install pandas seaborn
# pip freeze > requirements.txt
pip install -r requirements.txt

python main.py > ../../out/out-arrange-dataset.json 2> ../../out/out-arrange-dataset.json
