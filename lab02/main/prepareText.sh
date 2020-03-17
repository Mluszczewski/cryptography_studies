#!/bin/bash

cat orig.txt | tr -d "," | tr -d "." | tr -d " " | tr -d "’" | tr -d "'" | tr -d "-" | tr -d ";" | tr -d ":" | tr -d "!" | tr -d "?" | tr -d '"' | tr [:lower:] [:upper:]
awk  'length <= 1968 { printf "%-1968s\n",$0 }'  "orig.txt"

