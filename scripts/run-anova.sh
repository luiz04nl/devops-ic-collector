#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

# rm -rf out/anova/ || echo "Skipped"
rm out-anova.json || echo "Skipped"

mkdir -p out/anova/

cd src/anova

mkdir .env || echo "Skipped"
python -m venv .env || echo "Skipped"
. .env/bin/activate

# pip install pandas statsmodels
# pip freeze > requirements.txt
pip install -r requirements.txt

python main.py > ../../out/out-anova.json 2> ../../out/out-anova.json
