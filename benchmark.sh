#!/bin/sh

for i in $(seq 1 $2); do
    ./$1 > /dev/null
done
