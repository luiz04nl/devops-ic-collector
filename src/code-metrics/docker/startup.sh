#!/bin/bash
REPOSITORY=$1

if [ -z "$REPOSITORY" ]; then
  echo "Nenhum identificador de projeto fornecido. Uso: /usr/local/bin/startup.sh <REPOSITORY>"
  exit 1
fi

echo "The first argument is: $1"
echo "REPOSITORY: $1"

ls -la /app/ck-0.7.0-jar-with-dependencies.jar

PROJECT_DIR="/repos/$REPOSITORY"
PROJECT_OUT="/repos/$REPOSITORY/ck-out/"

mkdir -p $PROJECT_OUT

# OBS https://github.com/mauricioaniche/ck

java -jar /app/ck-0.7.0-jar-with-dependencies.jar  \
	$PROJECT_DIR \
	true \
	0 \
	true \
	$PROJECT_OUT \
	[]