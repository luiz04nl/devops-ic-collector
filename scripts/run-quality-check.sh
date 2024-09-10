#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

current_date_time="`date +%Y-%m-%d-%H-%M-%S`"

qualityCheckOutputDirectory="out/quality-check-repos"
if [ -d "$qualityCheckOutputDirectory" ] && [ "$(ls -A $qualityCheckOutputDirectory)" ]; then
    mkdir -p $qualityCheckOutputDirectory-bkp-$current_date_time/
    cp -rvf $qualityCheckOutputDirectory/* $qualityCheckOutputDirectory-bkp-$current_date_time/ | echo "$qualityCheckOutputDirectory not found"
    cat  out/out-quality-check-repos.json >> out/out-quality-check-repos-bkp-$current_date_time.json
    echo "\n###### Starting run-quality-check.sh ########\n" >> out/out-quality-check-repos-bkp.json
fi

qualityCheckOutputDirectory="out/quality-check-repos"
if [ -d "$qualityCheckOutputDirectory" ]; then
    rm -r $qualityCheckOutputDirectory
fi
mkdir -p $qualityCheckOutputDirectory

cd src/quality-check-repos
chmod +x *sh

docker ps

# mkdir -p ./docker/sonarqube/data
# mkdir -p ./docker/sonarqube/extensions
# mkdir -p ./docker/sonarqube/logs
# mkdir -p ./docker/postgres/postgresql
# mkdir -p ./docker/postgres/data
# docker-compose -f sonar-docker-compose.yml down
# docker-compose -f sonar-docker-compose.yml up -d
# sleep 20

go run . > ../../out/out-quality-check-repos.json 2> ../../out/out-quality-check-repos.json

TOTAL_TRIES=`ls ../../out/quality-check-repos/* | grep out.txt | wc -l`
echo "total tries $TOTAL_TRIES"

TRIES=`ls ../../out/quality-check-repos/* | grep out.txt | nl`
echo "tries $TRIES"
