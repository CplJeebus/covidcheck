#!/bin/bash

getupdate(){
    curl https://opendata.ecdc.europa.eu/covid19/casedistribution/json/ > today.json
    cat today.json | sed 's/Cumulative_number_for_14_days_of_COVID-19_cases_per_100000/c14d100k/g'  > today-fixed.json
}


DEBUG=FALSE

if ! [ -z "$(date +"%b %d" | sed  -n 's/0\([1-9]\)/\1/p' | tr -d '\n' )" ]
then
    TODAY=$(date +"%b %d" | sed  -n 's/0\([1-9]\)/\1/p' | tr -d '\n' )
else
    TODAY=$(date +"%b %d")
fi

TODAYF=$(ls -lt today.json | awk '{print $6,$7}')

log(){
    if [[ ${DEBUG} == "TRUE" ]]
    then
        echo $@ | od -a
    fi
}

cases(){
Cty=$1
cat today-fixed.json | jq --arg C $Cty -r '.records[] | select(.geoId == $C ) | "\(.c14d100k) \(.geoId) \(.dateRep)" ' \
        | head -n $NUM \
        | awk '{printf "%.2f\t%s\t%s\n", $1,$2,$3}'
}

deaths(){
Cty=$1
cat today-fixed.json | jq --arg C $Cty -r '.records[] | select(.geoId == $C ) | "\(.deaths) \(.geoId) \(.dateRep)" ' \
        | head -n $NUM \
        | awk '{printf "%d\t%s\t%s\n", $1,$2,$3}'
}

newcases(){
Cty=$1
cat today-fixed.json | jq --arg C $Cty -r '.records[] | select(.geoId == $C ) | "\(.cases) \(.geoId) \(.dateRep)" ' \
        | head -n $NUM \
        | awk '{printf "%d\t%s\t%s\n", $1,$2,$3}'
}

checkfile(){
if [ -f ./today.json ]
then
    if [[ "$TODAY" == "$TODAYF" ]]
    then
        log Have Todays file
    else
        log Old file
        getupdate
    fi
else
    log No file
    getupdate
fi
}

main(){
checkfile

case $1 in
    f|F)
        getupdate
        ;;
    d|D)
        export NUM=$2
        shift
        for C in $*
        do
            deaths $C
        done
        ;;
    n|N)
        export NUM=$2
        shift
        for C in $*
        do
            newcases $C
        done
        ;;
    *)
        export NUM=$1
        shift
        for C in $*
        do
        cases $C
        done
        ;;
esac

}

main $*

