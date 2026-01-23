#!/bin/bash

# Run Script
# This script executes the CLI application and outputs results to result.txt

echo "Running CLI application..."

# Execute the binary and redirect output to result.txt
# ./order-controller > scripts/result.txt 2>&1

go run ./cmd/main.go

if [ $? -eq 0 ]; then
    echo "CLI application execution completed"
    echo "Output saved to scripts/result.txt"
else
    echo "CLI application execution failed"
    exit 1
fi