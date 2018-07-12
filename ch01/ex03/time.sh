#!/bin/bash

echo "echo2測定"
time cat ./date.txt |xargs ./echo2|tail -1

echo "echo3測定"
time cat ./date.txt |xargs ./echo3|tail -1
