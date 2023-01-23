#!/usr/bin/env sh
set -e
# set -x

echo "=== Insert Test Data ==="

curl -X POST -s localhost:8080 -d \
    '{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

curl -X POST -s localhost:8080 -d \
    '{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

curl -X POST -s localhost:8080 -d \
    '{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'

echo "=== Test Descriptions ==="

curl -X GET -s localhost:8080/0 | grep -q 'christmas eve bike class'
curl -X GET -s localhost:8080/1 | grep -q 'cross country skiing'
curl -X GET -s localhost:8080/2 | grep -q 'sledding'

echo "Success"
