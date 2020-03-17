#!/bin/bash

cat orig.txt | tr -d "," | tr -d "." | tr -d "’" | tr -d "'" | tr -d "-" | tr -d ";" | tr -d ":" | tr -d "!" | tr -d "?" | tr -d '"' | tr [:upper:] [:lower:] | awk '{printf("%-40s"((length($0)>40)?"":"")"\n", substr($0,1,40))}'
