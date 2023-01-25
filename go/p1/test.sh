#!/usr/bin/env sh
set -e
# set -x
grpcurl -plaintext localhost:8080 describe

echo "=== Insert Test Data ==="

grpcurl -plaintext -d '{ "description": "christmas eve bike class" }' localhost:8080 api.v1.Activity_Log/Insert

echo "=== Test Retrieve Descriptions ==="

grpcurl -plaintext -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve | grep -q 'christmas eve bike class'

echo "=== Test List ==="

grpcurl -plaintext localhost:8080 api.v1.Activity_Log/List | jq '.activities | length' | grep -q '1'

echo "Success"
