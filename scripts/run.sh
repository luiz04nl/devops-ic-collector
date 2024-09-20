#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

current_date_time="`date +%Y-%m-%d-%H-%M-%S`";

### Data Extraction Part 1

sh ./scripts/run-create-dataset.sh &&

sh ./scripts/run-clone-repos.sh &&

sh ./scripts/run-check-devops-and-tools.sh &&

sh ./scripts/run-check-agile-and-behaviors.sh  &&

sh ./scripts/run-contributors.sh &&

sh ./scripts/run-prepare-quality-check.sh &&

sh ./scripts/run-quality-check.sh &&

### Data Analysis

# sh ./scripts/run-arrange-dataset.sh &&
## Not Used - Categories - out/arrange-dataset/CategoryGroup-Quantity.png
## Figure 1 - Mean of Bugs by Group - out/arrange-dataset/CategoryGroup-Bugs.png
## Figure 2 - Mean of Vulnerabilities by Group - out/arrange-dataset/CategoryGroup-Vulnerabilities.png
## Figure 3 - Mean of Security Rating by Group - out/arrange-dataset/CategoryGroup-SecurityRating.png
## Figure 4 - Mean of Code Smells by Group - out/arrange-dataset/CategoryGroup-CodeSmells.png

# sh ./scripts/run-anova.sh &&
# sh ./scripts/run-regression.sh &&
# sh ./scripts/run-chi-square.sh &&
# sh ./scripts/run-contributors-regression.sh &&

echo "Run finished"
