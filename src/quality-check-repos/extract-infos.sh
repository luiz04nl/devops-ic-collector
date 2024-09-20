#!/usr/bin/env bash

REPOSITORY=$1
WITH_BUILD=${2:-true}
$SCRIPT_BASE_PATH=$3

. ./build-and-scan.sh $REPOSITORY $WITH_BUILD $SCRIPT_BASE_PATH

export projectType=$projectType
export projectTypeVersion=$projectTypeVersion
export analysisSuccess=$analysisSuccess
export filesAtRootDir=$filesAtRootDir

pageSize=500

echo "######### Geting Issues ###################"
currentPage=1
totalPages=1 # apenas para executar a primeira vez
while [ $currentPage -le $totalPages ]
do
    response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
    if [ $(echo $response | jq '.errors') != "null" ]; then
        echo "Erro ao buscar dados: $(echo $response | jq '.errors')"
        exit 1
    fi
    echo $response | jq '.' > $SCRIPT_BASE_PATH/../../out/quality-check-repos/$REPOSITORY-ISSUES-page-${currentPage}.json

    total=$(echo $response | jq '.paging.total')
    totalPages=$(echo "($total + $pageSize - 1) / $pageSize" | bc)

    currentPage=$((currentPage + 1))
done

echo "######### Geting Code Smells ###################"
currentPage=1
totalPages=1 # apenas para executar a primeira vez
while [ $currentPage -le $totalPages ]
do
    response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&types=CODE_SMELL&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
    if [ $(echo $response | jq '.errors') != "null" ]; then
        echo "Erro ao buscar dados: $(echo $response | jq '.errors')"
        exit 1
    fi
    echo $response | jq '.' > $SCRIPT_BASE_PATH/../../out/quality-check-repos/$REPOSITORY-CODE_SMELL-page-${currentPage}.json

    total=$(echo $response | jq '.paging.total')
    totalPages=$(echo "($total + $pageSize - 1) / $pageSize" | bc)

    currentPage=$((currentPage + 1))
done
echo "############################"

# 1. Severidade
# SonarQube atribui uma severidade a cada code smell, o que ajuda a determinar a importância de resolver o problema. As severidades incluem:
# Blocker: Problemas críticos que podem causar erros graves.
# Critical: Problemas sérios que podem afetar significativamente o funcionamento ou a manutenção do código.
# Major: Problemas que devem ser corrigidos, mas que não são considerados críticos.
# Minor: Problemas de menor impacto que podem melhorar a legibilidade ou a manutenção do código.
# Info: Questões informativas que não precisam necessariamente de ação, mas podem ser úteis para melhorias.

# 2. Categorias de Code Smells
# Os code smells no SonarQube são agrupados em várias categorias, cada uma correspondendo a diferentes tipos de problemas:
# Duplicações: Código duplicado que pode ser consolidado para evitar redundância.
# Design: Problemas relacionados ao design do código, como complexidade excessiva ou violação de princípios de design.
# Manutenção: Problemas que dificultam a manutenção do código, como funções muito longas ou variáveis mal nomeadas.
# Leitura: Questões que afetam a legibilidade do código, como convenções de nomeação e formatação inadequada.
# Confiabilidade: Problemas que podem causar falhas no sistema.
# Segurança: Vulnerabilidades de segurança que devem ser corrigidas.

requestSonarMeasures() {
  METRIC=$1
  curl -s "${SONAR_URL}/api/measures/component?component=${REPOSITORY}&metricKeys=$METRIC" \
--header "Authorization: Basic $ENCODED_CREDENTIALS" \
> $SCRIPT_BASE_PATH/../../out/quality-check-repos/$REPOSITORY-$METRIC.json
}

requestSonarMeasures "sqale_rating"
#A=0-0.05, B=0.06-0.1, C=0.11-0.20, D=0.21-0.5, E=0.51-1

requestSonarMeasures "reliability_rating"
# A = 0 Bugs
# B = at least 1 Minor Bug
# C = at least 1 Major Bug
# D = at least 1 Critical Bug
# E = at least 1 Blocker Bug

requestSonarMeasures "complexity"
requestSonarMeasures "cognitive_complexity"
requestSonarMeasures "duplicated_blocks"
requestSonarMeasures "duplicated_files"
requestSonarMeasures "duplicated_lines"
requestSonarMeasures "code_smells"
requestSonarMeasures "ncloc"
requestSonarMeasures "sqale_index"
requestSonarMeasures "sqale_debt_ratio"
requestSonarMeasures "quality_gate_details"
requestSonarMeasures "bugs"
requestSonarMeasures "vulnerabilities"
requestSonarMeasures "security_rating"
requestSonarMeasures "classes"
requestSonarMeasures "comment_lines"
requestSonarMeasures "coverage"
requestSonarMeasures "tests"

json=$(jq -n \
  --arg projectType "$projectType" \
  --arg projectTypeVersion "$projectTypeVersion" \
  --arg REPOSITORY "$REPOSITORY" \
  --arg analysisSuccess "$analysisSuccess" \
  --arg filesAtRootDir "$filesAtRootDir" \
  '{projectType: $projectType, projectTypeVersion: $projectTypeVersion, repository: $REPOSITORY, analysisSuccess, $analysisSuccess, filesAtRootDir: $filesAtRootDir'})
echo  $json > $SCRIPT_BASE_PATH/../../out/quality-check-repos/$REPOSITORY.json

export extractInfosSuccess=1