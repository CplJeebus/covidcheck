#!/bin/bash

i=0
j=1
cat us-population-data.json | jq -r .[].State | while read st
do
    code=$(grep "$st" us-state-codes.txt | cut -d ',' -f2 | head -n 1)
    cat pop-${i} | sed "s/\"State\": \"${st}\",/\"State\": \"${st}\",\
\"Code\": \"${code}\",/" > pop-${j} 
    j=$((j+1))
    i=$((i+1))
done

