#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
./build/a --add "overhead press: 70lbs"
./build/a --add "20 minute walk"

echo "=== Retrieve Records ==="
./build/a --get 0 | grep "overhead press"
./build/a --get 1 | grep "20 minute walk"
