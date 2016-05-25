#! /bin/bash

interval=$1

while true
do
    ave=0
    before=`wc -l /root/papertest/generator/ttl | cut -f1 -d " "`
    sleep $interval 
    now=`wc -l /root/papertest/generator/ttl | cut -f1 -d " "`
    gap=$((now - before))
    ave=`tail -$gap /root/papertest/generator/ttl | awk '{sum += $1} END{print sum/NR}'`
    echo $ave | grep [0-9] > /dev/null
    if [ $? -ne 0 ]
    then
        ave=0.0
    fi
    echo "${gap}'s average: ${ave}"
done

