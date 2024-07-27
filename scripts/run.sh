#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

# current_date_time="`date +%Y-%m-%d-%H-%M-%S`";

# rm out/*json || true
# mv ./database/sqlite/repository-dataset.db ./database/sqlite/bkp/repository-dataset-$current_date_time.db

# sh ./scripts/run-create-dataset.sh &&

# sh ./scripts/run-clone-repos.sh &&

# sh ./scripts/run-check-devops-and-tools.sh &&

# sh ./scripts/run-check-agile-and-behaviors.sh  &&

# sh ./scripts/run-quality-check.sh  &&

# echo "Run finished"
