#!/bin/zsh

YEAR_MONTH="$(date '+%y.%m')"
BUILD="$(( ($(date +%s) - $(date -d'2000/1/1' +%s) ) / 60 / 60 / 24 ))"
REVISION="$(( ( $(date +%s) - $(date "-d$(date +%Y/%m/%d)" +%s) ) / 2 ))"
VERSION="${YEAR_MONTH}.${BUILD}.${REVISION}"

echo "${VERSION}"
