#!/bin/bash

n=$1

if [ "$n" -gt 100 ]; then
    n=100
fi

for ((i=1; i<=n; i++)); do
    echo "This is loop number $i"
done